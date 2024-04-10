package main

import (
	"context"
	"flag"
	"log"

	"github.com/fhke/infrastructure-abstraction/dsl"
	"github.com/fhke/infrastructure-abstraction/sdk/client"
)

func main() {
	var (
		srcDir        = flag.String("source-dir", ".", "Source directory to use for Terraform")
		serverBaseURL = flag.String("server-base-url", "http://localhost:9001", "Base URL for API server")
		stackName     = flag.String("stack-name", "", "Name of Stack to deploy")
	)
	flag.Parse()
	if *stackName == "" {
		log.Fatal("Flag -stack-name must be set")
	}

	if err := dsl.Run(
		context.TODO(),
		client.New(*serverBaseURL),
		*srcDir,
		*stackName,
	); err != nil {
		panic(err)
	}
}
