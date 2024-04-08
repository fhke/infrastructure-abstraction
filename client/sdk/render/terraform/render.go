package terraform

import (
	"context"
	"fmt"

	"github.com/fhke/infrastructure-abstraction/sdk/render"
	"github.com/samber/lo"
)

func (t *terraformRenderer) RenderStack(ctx context.Context, stackInput render.Stack) (TerraformStack, error) {
	stackModules, err := t.cl.BuildStack(
		ctx,
		stackInput.Name,
		stackInput.Repository,
		lo.MapToSlice(stackInput.Modules, func(_ string, mod render.Module) string {
			return mod.Name
		}),
	)
	if err != nil {
		return TerraformStack{}, fmt.Errorf("error building stack: %w", err)
	}

	tfStack := TerraformStack{
		Modules: make(map[string]Module),
	}

	for name, mod := range stackInput.Modules {
		stackModule, ok := stackModules.Modules[mod.Name]
		if !ok {
			return TerraformStack{}, fmt.Errorf("response from API does not contain module %s", mod.Name)
		}

		tfStack.Modules[name] = Module{
			Source:  stackModule.Source,
			Version: stackModule.Version,
			Inputs:  mod.Inputs,
		}
	}

	return tfStack, nil
}
