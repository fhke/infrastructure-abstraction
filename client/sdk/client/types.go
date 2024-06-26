package client

type (
	Stack struct {
		Modules map[string]BuildStackOutModule `json:"modules"`
	}
	BuildStackOutModule struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	}
)
