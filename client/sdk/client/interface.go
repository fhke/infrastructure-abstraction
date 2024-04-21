package client

import (
	"context"

	"github.com/Masterminds/semver"
)

type Client interface {
	CreateModuleVersion(ctx context.Context, name, source string, version *semver.Version) error
	BuildStack(ctx context.Context, name, repository string, moduleNames []string) (Stack, error)
	GetModuleNames(ctx context.Context) ([]string, error)
	SetStackModules(ctx context.Context, name, repository string, moduleVersions map[string]string) (Stack, error)
}
