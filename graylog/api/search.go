package api

import (
	"strconv"
	"time"
)

type TimeRangeApi struct {
	Type  string    `json:"type"`
	Range int       `json:"range"`
	From  *time.Time `json:"from,omitempty"`
	To    *time.Time `json:"to,omitempty"`
}

func (t *TimeRangeApi) MarshalJSON() ([]byte, error) {
	if t.Type == "absolute" && t.From != nil && t.To != nil {
		return []byte(`{"type": "` + t.Type + `", "from":"` + t.From.Format(time.RFC3339) +`", "to":"`+ t.To.Format(time.RFC3339) +`"}`), nil
	}
	return []byte(`{"type": "` + t.Type + `", "range":"` + strconv.Itoa(t.Range) + `"}`), nil
}

type InnerQueryApi struct {
	Type        string `json:"type"`
	QueryString string `json:"query_string"`
}

type SortApi struct {
	Type string `json:"type,omitempty"`
	Field string `json:"field"`
	Order string `json:"order,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type SeriesApi struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Field string `json:"field"`
}

type InternalApi struct {
	Type    string `json:"type"`
	Scaling int    `json:"scaling"`
}

type GroupApi struct {
	Type     string       `json:"type"`
	Field    string       `json:"field"`
	Internal *InternalApi `json:"internal"`
}

type InnerFilterApi struct {
	Type    string   `json:"type"`
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Filters *[]string `json:"filters"`
}

type FilterApi struct {
	Type    string            `json:"type"`
	Filters *[]InnerFilterApi `json:"filters"`
}

type SearchTypeApi struct {
	Id           string         `json:"id"`
	Timerange    *TimeRangeApi  `json:"timerange"`
	Query        *InnerQueryApi `json:"query"`
	Streams      []string       `json:"streams"`
	Name         string       `json:"name"`
	Limit        int          `json:"limit"`
	Offset       int          `json:"offset"`
	Sort         []*SortApi   `json:"sort"`
	Decorators   []string     `json:"decorators"`
	Type         string       `json:"type"`
	Series       []*SeriesApi `json:"series,omitempty"`
	Rollup       bool         `json:"rollup,omitempty"`
	RowGroups    []*GroupApi  `json:"row_groups,omitempty"`
	ColumnGroups []*GroupApi  `json:"column_groups,omitempty"`
	Filter       *FilterApi   `json:"filter,omitempty"`
}

type OuterQueryApi struct {
	Id          string           `json:"id"`
	Timerange   *TimeRangeApi    `json:"timerange"`
	Query       *InnerQueryApi   `json:"query"`
	Filter     *FilterApi        `json:"filter,omitempty"`
	SearchTypes []*SearchTypeApi `json:"search_types"`
}

type SearchApi struct {
	Id string                `json:"id"`
	Queries []*OuterQueryApi `json:"queries"`
}
