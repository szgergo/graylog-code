package template

type StreamRuleTemplate struct {
	Id          string
	Field       string `yaml:"field"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Inverted    bool   `yaml:"inverted"`
	Value       string `yaml:"value"`
}

func (srt *StreamRuleTemplate) Equals(other StreamRuleTemplate) bool {
	return srt.Description == other.Description &&
		srt.Field == other.Field &&
		srt.Type == other.Type &&
		srt.Inverted == other.Inverted &&
		srt.Value == other.Value
}

type StreamTemplate struct {
	Id string
	MatchingType string `yaml:"matching_type"`
	Description string `yaml:"description"`
	Title string `yaml:"title"`
	ContentPack *interface{} `yaml:"content_pack,omitempty"`
	RemoveMatchesFromDefaultStream bool `yaml:"remove_matches_from_default_stream"`
	IndexSetName string `yaml:"index_set_name"`
	AutoStart bool `yaml:"auto_start"`
	Rules *[]StreamRuleTemplate `yaml:"rules,omitempty"`
}

func StreamRuleTemplateArrayEquals(a []StreamRuleTemplate, b[]StreamRuleTemplate) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		isEqual := false
		for j := range b {
			if a[i].Equals(b[j]) {
				isEqual = true
			}
		}
		if !isEqual {
			return false
		}
	}
	return true
}

func (st *StreamTemplate) Equals(other StreamTemplate) bool {
	return st.MatchingType == other.MatchingType &&
		st.Description == other.Description &&
		st.Title == other.Title &&
		st.RemoveMatchesFromDefaultStream == other.RemoveMatchesFromDefaultStream &&
		st.IndexSetName == other.IndexSetName &&
		StreamRuleTemplateArrayEquals(*st.Rules,*other.Rules)
}

/*
From: http://web.mnstate.edu/peil/MDEV102/U1/S6/Complement3.htm

Complement of a Set of Streams: The complement of a set of Streams, denoted diff,
is the set of all elements in the given universal set a that are not in b.
diff = {x ∈ a : x ∉ b}.
*/
func StreamComplement(a, b []StreamTemplate) (diff []StreamTemplate) {

	if len(b) == 0 {
		return a
	}

	for i := range a {
		for j := range b {
			if a[i].Equals(b[j]) {
				break
			} else if j + 1 == len(b) {
				diff = append(diff,a[i])
			}
		}
	}
	return diff
}
