package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fhke/infrastructure-abstraction/client/util"
)

func (r *restyClient) GetModuleNames(ctx context.Context) ([]string, error) {
	type response struct {
		Names []string `json:"names"`
	}

	resp, err := r.cl.R().SetContext(ctx).Get("/api/modules")
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	} else if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	out, err := util.UnmarshalJSONRestyResp[response](resp)
	if err != nil {
		return nil, fmt.Errorf("error decoding body: %w", err)
	}

	return out.Names, nil
}
