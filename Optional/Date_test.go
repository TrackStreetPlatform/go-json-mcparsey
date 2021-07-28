package Optional

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue time.Time
	}
	strTime := "1997-03-01T18:45:26Z"
	testTime, _ := time.Parse("2006-01-02T15:04:05Z", strTime)
	testDateTime := primitive.NewDateTimeFromTime(testTime)
	tests := []struct {
		name   string
		input  inputStruct
		output time.Time
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testTime},
				Key:          "NonExisting",
				DefaultValue: time.Time{},
			},
			output: time.Time{},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": strTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: testTime,
		},
		{
			name: "CaseStringLayoutError",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "2014-11-12T11:45:260Z"},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: time.Time{},
		},
		{
			name: "CaseTime",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: testTime,
		},
		{
			name: "CaseDateTime",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testDateTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: testDateTime.Time(),
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42},
				Key:          "value",
				DefaultValue: time.Time{}},
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
