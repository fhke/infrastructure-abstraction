package terraform

import (
	"fmt"

	"github.com/fhke/infrastructure-abstraction/sdk/client"
	"github.com/fhke/infrastructure-abstraction/sdk/deploy"
	tfExecutor "github.com/fhke/infrastructure-abstraction/sdk/executor/terraform"
	tfRender "github.com/fhke/infrastructure-abstraction/sdk/render/terraform"
)

func New(cl client.Client) (deploy.Deployer[tfRender.TerraformStack], error) {
	ex, err := tfExecutor.New(".generatedTerraform")
	if err != nil {
		return nil, fmt.Errorf("error configuring executor: %w", err)
	}

	return deploy.New[tfRender.TerraformStack](ex, tfRender.New(cl)), nil
}
