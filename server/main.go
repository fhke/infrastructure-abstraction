package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fhke/infrastructure-abstraction/server/controller"
	"github.com/fhke/infrastructure-abstraction/server/handler"
	modulesRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository/file"
	stacksRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository/file"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	var (
		dataPathStacks  = flag.String("data-stacks", "./tmp/stacks.json", "Path to file containing stacks database.")
		dataPathModules = flag.String("data-modules", "./tmp/modules.json", "Path to file containing modules database.")
		listenAddr      = flag.String("listen-addr", "127.0.0.1:9001", "Address to bind server to")
		debug           = flag.Bool("debug", false, "Enable debug logging")
	)
	flag.Parse()

	log := mustLog(*debug)

	stacksDataFile := lo.Must(os.OpenFile(*dataPathStacks, os.O_CREATE|os.O_RDWR, 0644))
	modulesDataFile := lo.Must(os.OpenFile(*dataPathModules, os.O_CREATE|os.O_RDWR, 0644))

	modRepo := lo.Must(modulesRepository.New(modulesDataFile))
	stRepo := lo.Must(stacksRepository.New(stacksDataFile))

	ctrl := controller.New(log.Named("controller"), stRepo, modRepo)

	h := handler.New(log.Named("handler"), ctrl)

	gin.SetMode(gin.ReleaseMode)
	ginE := gin.New()
	ginE.Use(
		ginzap.Ginzap(log.Desugar().Named("request"), time.RFC3339, true),
		ginzap.RecoveryWithZap(log.Desugar().Named("recovery"), true),
	)
	ginE.POST("/api/stack/build", h.HandleBuildStack)
	ginE.POST("/api/modules", h.HandleCreateModuleVersion)

	srv := &http.Server{
		Addr:    *listenAddr,
		Handler: ginE,
	}

	killCh := make(chan os.Signal, 1)
	signal.Notify(killCh, os.Interrupt)

	go func() {
		<-killCh
		log.Info("Stopping server in response to interrupt")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Panicw("Error closing server", "error", err)
		}
		log.Info("Server shutdown successfully")
	}()

	log.Infow("Starting server", "addr", *listenAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicw("Server listen error", "error", err)
	}
	log.Info("Exiting")
}

func mustLog(debug bool) *zap.SugaredLogger {
	cnf := zap.NewProductionConfig()
	cnf.Sampling = nil
	if debug {
		cnf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	return lo.Must(cnf.Build()).Sugar()
}
