package cdk

import "fmt"

type GenericModuleBuilder struct {
	name   string
	module string
	inputs map[string]any
}

func newGenericModuleBuilder(name, module string) *GenericModuleBuilder {
	return &GenericModuleBuilder{
		name:   name,
		module: module,
		inputs: make(map[string]any),
	}
}

func (g *GenericModuleBuilder) WithInput(name string, value any) *GenericModuleBuilder {
	g.inputs[name] = value
	return g
}

func (g *GenericModuleBuilder) GetOutput(name string) string {
	return fmt.Sprintf("${%s}", g.GetOutputRef(name))
}

func (g *GenericModuleBuilder) GetOutputRef(name string) string {
	return fmt.Sprintf("module.%s.%s", g.name, name)
}

func (g *GenericModuleBuilder) Module() string {
	return g.module
}

func (g *GenericModuleBuilder) Inputs() map[string]any {
	return g.inputs
}
