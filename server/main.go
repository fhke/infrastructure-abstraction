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
	modulesRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
	dynamoModuleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository/dynamo"
	fileModuleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository/file"
	stackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
	dynamoStackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository/dynamo"
	fileStackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository/file"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	var (
		// Storage configuration flags.
		storageStrategy    = flag.String("storage", "file", "Data storage strategy to use. Must be either \"dynamo\" or \"file\".")
		dataPathStacks     = flag.String("file.stacks", "./tmp/stacks.json", "Path to file containing stacks database. Only used for \"file\" storage strategy.")
		dataPathModules    = flag.String("file.modules", "./tmp/modules.json", "Path to file containing modules database. Only used for \"file\" storage strategy.")
		dynamoTableStacks  = flag.String("dynamo.stacks", "stacks", "Name of DynamoDB table for stacks data. Only used for \"dynamo\" storage strategy.")
		dynamoTableModules = flag.String("dynamo.modules", "modules", "Name of DynamoDB table for modules data. Only used for \"dynamo\" storage strategy.")

		// Server configuration flags.
		listenAddr = flag.String("listen-addr", "127.0.0.1:9001", "Address to bind server to")
		debug      = flag.Bool("debug", false, "Enable debug logging")

		// impl vars
		stRepo  stackRepository.Repository
		modRepo modulesRepository.Repository
	)
	flag.Parse()

	log := mustLog(*debug)

	switch ss := *storageStrategy; ss {
	case "dynamo":
		log.Info("Using dynamo storage strategy")
		cl := dynamodb.NewFromConfig(lo.Must(config.LoadDefaultConfig(context.TODO())))
		modRepo = dynamoModuleRepository.New(cl, *dynamoTableModules)
		stRepo = dynamoStackRepository.New(cl, *dynamoTableStacks)
	case "file":
		log.Info("Using file storage strategy")
		stacksDataFile := lo.Must(os.OpenFile(*dataPathStacks, os.O_CREATE|os.O_RDWR, 0644))
		modulesDataFile := lo.Must(os.OpenFile(*dataPathModules, os.O_CREATE|os.O_RDWR, 0644))
		modRepo = lo.Must(fileModuleRepository.New(modulesDataFile))
		stRepo = lo.Must(fileStackRepository.New(stacksDataFile))
	default:
		log.Panicf("Unknown storage strategy %q", ss)
	}

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
