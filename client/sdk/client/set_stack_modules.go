package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fhke/infrastructure-abstraction/client/util"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/request"
)

func (r *restyClient) SetStackModules(ctx context.Context, name, repository string, moduleVersions map[string]string) (Stack, error) {
	resp, err := r.cl.
		R().
		SetContext(ctx).
		SetBody(request.PatchStack{
			Name:           name,
			Repository:     repository,
			ModuleVersions: moduleVersions,
		}).
		Patch("/api/stack")
	if err != nil {
		return Stack{}, fmt.Errorf("request error: %w", err)
	} else if resp.StatusCode() != http.StatusCreated {
		return Stack{}, fmt.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode())
	}

	return util.UnmarshalJSONRestyResp[Stack](resp)
}
