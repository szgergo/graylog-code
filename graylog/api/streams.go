package api

type StreamApi struct {
	Id string `json:"id"`
	User string `json:"creator_user_id,omitempty"`
	MatchingType string `json:"matching_type"`
	Description string `json:"description"`
	Title string `json:"title"`
	ContentPack *interface{} `json:"content_pack,omitempty"`
	RemoveMatchesFromDefaultStream bool `json:"remove_matches_from_default_stream"`
	IndexSetId string `json:"index_set_id"`
	Rules *[]StreamRuleCreationApi `json:"rules,omitempty"`

}

type StreamsApi struct {
	Total int
	Streams *[]StreamApi
}
