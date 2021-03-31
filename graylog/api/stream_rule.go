package api

type StreamRuleCreationApi struct {
	Id string `json:"id,omitempty"`
	Field string `json:"field"`
	Description string `json:"description"`
	Type int `json:"type"`
	Inverted bool `json:"inverted"`
	Value string `json:"value"`
}