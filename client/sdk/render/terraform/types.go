package terraform

import "encoding/json"

type (
	TerraformStack struct {
		Modules map[string]Module `json:"module"`
	}
	Module struct {
		Source  string
		Version string
		Inputs  map[string]any
	}
)

var _ json.Marshaler = (TerraformStack{}).Modules[""]
