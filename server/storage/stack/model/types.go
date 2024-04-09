package model

type (
	Stack struct {
		Name       string            `json:"name" dynamodbav:"name"`
		Repository string            `json:"repository" dynamodbav:"repository"`
		Modules    map[string]Module `json:"modules" dynamodbav:"modules"`
	}

	Module struct {
		Source  string `json:"source" dynamodbav:"source"`
		Version string `json:"version" dynamodbav:"version"`
	}
)
