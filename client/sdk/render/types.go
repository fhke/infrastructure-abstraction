package render

import "context"

type (
	Renderer[T any] interface {
		RenderStack(context.Context, Stack) (T, error)
	}
	Stack struct {
		Name       string
		Repository string
		Modules    map[string]Module
	}
	Module struct {
		Name   string
		Inputs map[string]any
	}
)
