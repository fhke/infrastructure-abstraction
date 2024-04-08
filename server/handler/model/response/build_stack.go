package response

type (
	BuildStack struct {
		Modules map[string]BuildStackModule `json:"modules"`
	}
	BuildStackModule struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	}
)
