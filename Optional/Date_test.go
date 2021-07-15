package Optional

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	strTime := "1997-03-01T18:45:26Z"
	testTime, _ := time.Parse("2006-01-02T15:04:05Z", strTime)
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue time.Time
		}
		output time.Time
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue time.Time
			}{
				Origin: map[string]interface{}{
					"value": testTime,
				}, Key: "NonExisting", DefaultValue: time.Time{}},
			output: time.Time{},
		},
		{
			name: "CaseString",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue time.Time
			}{
				Origin: map[string]interface{}{
					"value": strTime,
				}, Key: "value", DefaultValue: time.Time{}},
			output: testTime,
		},
		{
			name: "CaseStringLayoutError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue time.Time
			}{
				Origin: map[string]interface{}{
					"value": "2014-11-12T11:45:260Z",
				}, Key: "value", DefaultValue: time.Time{}},
			output: time.Time{},
		},
		{
			name: "CaseTime",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue time.Time
			}{
				Origin: map[string]interface{}{
					"value": testTime,
				}, Key: "value", DefaultValue: time.Time{}},
			output: testTime,
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue time.Time
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "value", DefaultValue: time.Time{}},
			output: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Date(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected Date(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output,
					got,
				)
			}
		})
	}
}
