package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fhke/infrastructure-abstraction/client/util"
)

func (c *restyClient) BuildStack(ctx context.Context, name, repository string, moduleNames []string) (Stack, error) {
	type request struct {
		Name        string   `json:"name"`
		Repository  string   `json:"repository"`
		ModuleNames []string `json:"moduleNames"`
	}
	resp, err := c.cl.
		R().
		SetBody(request{
			Name:        name,
			Repository:  repository,
			ModuleNames: moduleNames,
		}).
		SetContext(ctx).
		Post("/api/stack/build")

	if err != nil {
		return Stack{}, fmt.Errorf("request error: %w", err)
	} else if resp.StatusCode() != http.StatusOK {
		return Stack{}, fmt.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	bso, err := util.UnmarshalJSONRestyResp[Stack](resp)
	if err != nil {
		return Stack{}, fmt.Errorf("error decoding body: %w", err)
	}

	return bso, nil
}
