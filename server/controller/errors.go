package controller

import (
	"fmt"

	"github.com/Masterminds/semver"
)

type ErrStackNotFound struct {
	Stack string
	Repo  string
}

func (e ErrStackNotFound) Error() string {
	return fmt.Sprintf("stack %s/%s not found", e.Repo, e.Stack)
}

type ErrModuleNotFound struct {
	Name string
}

func (e ErrModuleNotFound) Error() string {
	return fmt.Sprintf("module %s not found", e.Name)
}

type ErrModuleVersionAlreadyExists struct {
	Name    string
	Version *semver.Version
}

func (e ErrModuleVersionAlreadyExists) Error() string {
	return fmt.Sprintf("module %s already exists", e.Name)
}
