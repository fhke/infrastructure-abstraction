package cdk

type ModuleBuilder interface {
	Module() string
	Inputs() map[string]any
}
