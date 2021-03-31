package template

import "com/github/graylog-code/util"


type SavedSearchTemplate struct {
	Id 			string
	Name        string			   `yaml:"name,omitempty"`
	SearchQuery string			   `yaml:"search_query,omitempty"`
	Summary 	string             `yaml:"summary,omitempty"`
	Description string             `yaml:"description,omitempty"`
	Filters     *[]string          `yaml:"filters,omitempty"`
	TimeRange   *TimeRangeTemplate `yaml:"time_range,omitempty"`
	Users       *[]string          `yaml:"users,omitempty"`
}

func (s *SavedSearchTemplate) Equals(other SavedSearchTemplate) bool {
	return s.Name == other.Name &&
		s.SearchQuery == other.SearchQuery &&
		s.Description == other.Description &&
		s.Summary == other.Summary &&
		util.StringArrayEquals(*s.Filters, *other.Filters) &&
		timeRangePointerEquals(s.TimeRange, other.TimeRange) &&
		util.StringArrayEquals(*s.Users, *other.Users)
}

/*
From: http://web.mnstate.edu/peil/MDEV102/U1/S6/Complement3.htm

Complement of a Set of SavedSearches: The complement of a set of SavedSearches, denoted diff,
is the set of all elements in the given universal set a that are not in b.
diff = {x ∈ a : x ∉ b}.
*/
func SavedSearchComplement(a, b []SavedSearchTemplate) (diff []SavedSearchTemplate) {

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
