package terraform

import (
	"github.com/fhke/infrastructure-abstraction/sdk/client"
	"github.com/fhke/infrastructure-abstraction/sdk/render"
)

type terraformRenderer struct {
	cl client.Client
}

func New(cl client.Client) render.Renderer[TerraformStack] {
	return &terraformRenderer{
		cl: cl,
	}
}

var _ render.Renderer[TerraformStack] = new(terraformRenderer)
