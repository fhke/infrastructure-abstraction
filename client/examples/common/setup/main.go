package main

import (
	"context"
	"log"

	"github.com/Masterminds/semver"
	"github.com/fhke/infrastructure-abstraction/client/sdk/client"
)

func main() {
	cl := client.New("http://localhost:9001")
	mustCreateModuleVersion(cl, "aurora", "terraform-aws-modules/rds-aurora/aws", "9.3.1")
	mustCreateModuleVersion(cl, "aurora", "terraform-aws-modules/rds-aurora/aws", "9.1.0")
	mustCreateModuleVersion(cl, "aurora", "terraform-aws-modules/rds-aurora/aws", "8.5.0")
	mustCreateModuleVersion(cl, "vpc", "terraform-aws-modules/vpc/aws", "5.5.2")
	mustCreateModuleVersion(cl, "vpc", "terraform-aws-modules/vpc/aws", "5.7.1")
	mustCreateModuleVersion(cl, "eks", "terraform-aws-modules/eks/aws", "19.17.2")
	mustCreateModuleVersion(cl, "eks", "terraform-aws-modules/eks/aws", "20.8.4")
	mustCreateModuleVersion(cl, "eks", "terraform-aws-modules/eks/aws", "19.15.3")
}

func mustCreateModuleVersion(cl client.Client, name, source, version string) {
	semVersion := semver.MustParse(version)
	log.Printf("Creating module %q version %q with source %q", name, semVersion, source)
	if err := cl.CreateModuleVersion(context.TODO(), name, source, semVersion); err != nil {
		panic(err)
	}
}
