package api

type WidgetListApi struct {
	Widget map[string]string `json:"widget"`
}

type WidgetSettingsApi struct {
	SelectedFields      interface{}                   `json:"selected_fields"`
	StaticMessageListId interface{}                   `json:"static_message_list_id"`
	Titles              *WidgetListApi                `json:"titles"`
	Widgets             []*WidgetApi                  `json:"widgets"`
	WidgetMapping       *map[string][]string          `json:"widget_mapping"`
	Positions           map[string]*WidgetPositionApi `json:"positions"`
	Formatting          interface{}                   `json:"formatting"`
	DisplayModeSettings *DisplayModeSettingsApi       `json:"display_mode_settings,omitempty"`
}

type PivotConfigApi struct {
	Type    string      `json:"type"`
	Scaling interface{} `json:"scaling"`
}

type PivotApi struct {
	Field  string          `json:"field"`
	Type   string          `json:"type"`
	Config *PivotConfigApi `json:"config"`
}

type WidgetSeriesConfigApi struct {
	Name string `json:"name"`
}

type WidgetSeriesApi struct {
	Function string                 `json:"function"`
	Config   *WidgetSeriesConfigApi `json:"config"`
}

type WidgetConfigApi struct {
	RowPivots           []*PivotApi        `json:"row_pivots,omitempty"`
	ColumnPivots        []*PivotApi        `json:"column_pivots,omitempty"`
	Series              []*WidgetSeriesApi `json:"series,omitempty"`
	Sort                []*SortApi         `json:"sort"`
	Visualization       string             `json:"visualization,omitempty"`
	VisualizationConfig interface{}        `json:"visualization_config,omitempty"`
	FormattingSettings  interface{}        `json:"formatting_settings,omitempty"`
	Rollup              bool               `json:"rollup,omitempty"`
	EventAnnotation     bool               `json:"event_annotation,omitempty"`
	Fields              *[]string          `json:"fields"`
	ShowMessageRow      bool               `json:"show_message_row"`
	Decorators          *[]string          `json:"decorators"`
}

type WidgetApi struct {
	Id        string           `json:"id"`
	Type      string           `json:"type"`
	Filter    *FilterApi       `json:"filter"`
	Timerange *TimeRangeApi    `json:"timerange"`
	Query     *InnerQueryApi   `json:"query"`
	Streams   []string         `json:"streams"`
	Config    *WidgetConfigApi `json:"config"`
}

type WidgetPositionApi struct {
	Col    int    `json:"col"`
	Row    int    `json:"row"`
	Height int    `json:"height"`
	Width  string `json:"width"`
}

type ViewApi struct {
	Id 					string                      `json:"id,omitempty"`
	Type                string                        `json:"type"`
	Title               string                        `json:"title"`
	Summary             string                        `json:"summary"`
	Description         string                        `json:"description"`
	SearchId            string                        `json:"search_id"`
	Properties          []string                      `json:"properties"`
	Requires            interface{}                   `json:"requires,omitempty"`
	State               map[string]*WidgetSettingsApi `json:"state"`
	Owner string `json:"owner,omitempty"`
}

type DisplayModeSettingsApi struct {
	Positions map[string]WidgetPositionApi `json:"positions"`
}
