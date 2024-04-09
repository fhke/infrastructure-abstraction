package model

type (
	ModuleVersion struct {
		Name    string `json:"name" dynamodbav:"name"`       // name is the user-facing name of the module.
		Source  string `json:"source" dynamodbav:"source"`   // source is the terraform-readable source of the module.
		Version string `json:"version" dynamodbav:"version"` // version is the version of the module.
	}
)
