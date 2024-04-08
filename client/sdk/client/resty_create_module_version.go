package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Masterminds/semver"
)

func (c *restyClient) CreateModuleVersion(ctx context.Context, name, source string, version *semver.Version) error {
	type request struct {
		Name    string          `json:"name"`
		Source  string          `json:"source"`
		Version *semver.Version `json:"version"`
	}
	resp, err := c.cl.
		R().
		SetBody(request{
			Name:    name,
			Source:  source,
			Version: version,
		}).
		SetContext(ctx).
		Post("/api/modules")

	if err != nil {
		return fmt.Errorf("request error: %w", err)
	} else if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode())
	}

	return nil
}
