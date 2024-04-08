package terraform

import "encoding/json"

func (m Module) MarshalJSON() ([]byte, error) {
	marshalMap := make(map[string]any)
	for k, v := range m.Inputs {
		marshalMap[k] = v
	}
	marshalMap["source"] = m.Source
	marshalMap["version"] = m.Version

	return json.Marshal(marshalMap)
}
