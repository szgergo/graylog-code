package template

import (
	errorChecker "com/github/graylog-code/error"
	"time"
)

type TimeRangeTemplate struct {
	Type  string		`yaml:"type,omitempty"`
	Range int			`yaml:"range,omitempty"`
	From  *time.Time	`yaml:"from,omitempty"`
	To    *time.Time	`yaml:"to,omitempty"`
}

func (t TimeRangeTemplate) Equals(other TimeRangeTemplate) bool {
	return t.Type == other.Type &&
		t.Range == other.Range &&
		timePointerEquals(t.From, other.From) &&
		timePointerEquals(t.To, other.To)
}

func timePointerEquals(a *time.Time, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	if a != nil && b!= nil {
		return a.Equal(*b)
	}
	return false
}

func (t *TimeRangeTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp map[string]interface{}
	unmarshalErr := unmarshal(&tmp)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	var _type = tmp["type"].(string)
	var _range = 0
	if tmp["range"] != nil {
		_range = tmp["range"].(int)
	}
	var from, to string
	if tmp["from"] != nil && tmp["to"] != nil {
		from = tmp["from"].(string)
		to = tmp["to"].(string)
	}

	var tmp2 = TimeRangeTemplate{}
	tmp2.Type = _type
	tmp2.Range = _range
	if from != "" && to != "" {
		_from, fromTimeParseErr := time.Parse(time.RFC3339,from)

		errorChecker.Check(fromTimeParseErr)
		_to, toTimeParseErr := time.Parse(time.RFC3339,to)
		errorChecker.Check(toTimeParseErr)
		tmp2.From = &_from
		tmp2.To = &_to
	}

	*t = tmp2
	return nil
}

func timeRangePointerEquals(a *TimeRangeTemplate, b *TimeRangeTemplate) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(*b)
}
