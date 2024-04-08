package util

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func UnmarshalJSON[T any](data []byte) (T, error) {
	out := new(T)
	if err := json.Unmarshal(data, out); err != nil {
		return *out, err
	}
	return *out, nil
}

func UnmarshalJSONRestyResp[T any](resp *resty.Response) (T, error) {
	return UnmarshalJSON[T](resp.Body())
}
