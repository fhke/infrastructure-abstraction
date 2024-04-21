package request

type PatchStack struct {
	Name           string            `json:"name"`
	Repository     string            `json:"repository"`
	ModuleVersions map[string]string `json:"moduleVersions"`
}
