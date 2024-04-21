package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fhke/infrastructure-abstraction/server/controller"
	"github.com/fhke/infrastructure-abstraction/server/handler"
	dynamoModuleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository/dynamo"
	dynamoStackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository/dynamo"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	var (
		// Storage configuration flags.
		dynamoTableStacks  = flag.String("dynamo-table.stacks", "stacks", "Name of DynamoDB table for stacks data.")
		dynamoTableModules = flag.String("dynamo-table.modules", "modules", "Name of DynamoDB table for modules data.")

		// Server configuration flags.
		listenAddr = flag.String("listen-addr", "127.0.0.1:9001", "Address to bind server to")
		debug      = flag.Bool("debug", false, "Enable debug logging")
	)
	flag.Parse()

	log := mustLog(*debug)

	cl := dynamodb.NewFromConfig(lo.Must(config.LoadDefaultConfig(context.TODO())))

	h := handler.New(log.Named("handler"), controller.New(
		log.Named("controller"),
		dynamoStackRepository.New(cl, *dynamoTableStacks),
		dynamoModuleRepository.New(cl, *dynamoTableModules),
	))

	gin.SetMode(gin.ReleaseMode)
	ginE := gin.New()
	ginE.Use(
		ginzap.Ginzap(log.Desugar().Named("request"), time.RFC3339, true),
		ginzap.RecoveryWithZap(log.Desugar().Named("recovery"), true),
	)
	h.Register(ginE)

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
