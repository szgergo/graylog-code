package api

type StreamCreationApi struct {
	MatchingType string `json:"matching_type"`
	Description string `json:"description"`
	Title string `json:"title"`
	ContentPack *interface{} `json:"content_pack,omitempty"`
	RemoveMatchesFromDefaultStream bool `json:"remove_matches_from_default_stream"`
	IndexSetId string `json:"index_set_id"`
}

type StreamCreationResponseApi struct {
	Id string `json:"stream_id"`
}

