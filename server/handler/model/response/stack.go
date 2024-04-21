package response

type (
	Stack struct {
		Modules map[string]StackModule `json:"modules"`
	}
	StackModule struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	}
)
