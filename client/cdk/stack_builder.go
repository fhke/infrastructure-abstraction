package cdk

import (
	"context"
	"errors"
	"fmt"

	"github.com/fhke/infrastructure-abstraction/sdk/deploy"
	"github.com/fhke/infrastructure-abstraction/sdk/render"
	"github.com/fhke/infrastructure-abstraction/util/git"
	"github.com/samber/lo"
)

var (
	ErrNoModules = errors.New("stack must contain at least one module")
)

type StackBuilder[R any] struct {
	name       string
	repository string
	modules    map[string]ModuleBuilder
	deployer   deploy.Deployer[R]
}

func NewStackBuilder[R any](name string, deployer deploy.Deployer[R]) *StackBuilder[R] {
	return &StackBuilder[R]{
		name:     name,
		modules:  make(map[string]ModuleBuilder, 0),
		deployer: deployer,
	}
}

// SetRepoName overrides the derived repo name with a specific value.
func (s *StackBuilder[R]) SetRepoName(name string) *StackBuilder[R] {
	s.repository = name
	return s
}

func (s *StackBuilder[R]) Deploy(ctx context.Context) error {
	if len(s.modules) == 0 {
		return ErrNoModules
	}

	if s.repository == "" {
		// get repo name from config
		repoName, err := git.GetCurrentRepoSlug()
		if err != nil {
			return fmt.Errorf("error deriving git repository name: %w", err)
		}
		s.repository = repoName
	}

	err := s.deployer.Deploy(ctx, render.Stack{
		Name:       s.name,
		Repository: s.repository,
		Modules: lo.MapValues(s.modules, func(mod ModuleBuilder, _ string) render.Module {
			return render.Module{
				Name:   mod.Module(),
				Inputs: mod.Inputs(),
			}
		}),
	})
	if err != nil {
		return fmt.Errorf("error running stack: %w", err)
	}

	return nil
}

func (s *StackBuilder[R]) AddModuleBuilder(name string, builder ModuleBuilder) *StackBuilder[R] {
	s.modules[name] = builder
	return s
}
