package client

import (
	"context"

	"github.com/Masterminds/semver"
)

type Client interface {
	CreateModuleVersion(ctx context.Context, name, source string, version *semver.Version) error
	BuildStack(ctx context.Context, name, repository string, moduleNames []string) (BuildStackOut, error)
}
