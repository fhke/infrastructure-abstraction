package main

import (
	"context"

	"github.com/fhke/infrastructure-abstraction/cdk"
	"github.com/fhke/infrastructure-abstraction/sdk/client"
	terraformRunner "github.com/fhke/infrastructure-abstraction/sdk/deploy/terraform"
	"github.com/fhke/infrastructure-abstraction/sdk/render/terraform"
)

func main() {
	// configure clients
	cl := client.New("http://localhost:9001")
	rn, err := terraformRunner.New(cl)
	if err != nil {
		panic(err)
	}

	// create stack
	stack := cdk.NewStackBuilder[terraform.TerraformStack]("infra-test", rn)

	// Networking configuration
	vpcMod := stack.
		WithVPC("network").
		WithName("test").
		WithAZs(
			"eu-west-1a",
			"eu-west-1b",
			"eu-west-1c",
		).
		WithCIDR("10.0.0.0/16").
		WithDatabaseSubnets(
			"10.0.0.0/20",
			"10.0.16.0/20",
			"10.0.32.0/20",
		).
		WithPrivateSubnets(
			"10.0.64.0/20",
			"10.0.80.0/20",
			"10.0.96.0/20",
		).
		WithPublicSubnets(
			"10.0.128.0/20",
			"10.0.144.0/20",
			"10.0.160.0/20",
		)

	// Compute configuration
	stack.
		WithEKS("cluster").
		InVPC(vpcMod).
		WithClusterEndpointPrivateAccess(true).
		WithClusterEndpointPublicAccess(false).
		WithClusterName("test-apps")

	// Database configuration
	stack.
		WithAurora("database").
		InVPC(vpcMod).
		WithName("test-pg96").
		WithMasterUsername("root").
		WithEngine("aurora-postgresql").
		WithEngineVersion("14.5")

	// Deploy
	if err := stack.Deploy(context.TODO()); err != nil {
		panic(err)
	}
}
