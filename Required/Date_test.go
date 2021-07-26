package Required

import (
	"fmt"
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
	type outputStruct struct {
		Value         time.Time
		IsValid       bool
		MissingFields []string
	}
	strTime := "1997-03-01T18:45:26Z"
	testTime, _ := time.Parse("2006-01-02T15:04:05Z", strTime)
	testDateTime := primitive.NewDateTimeFromTime(testTime)
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testTime},
				Key:          "NonExisting",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         time.Time{},
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": strTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         testTime,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseStringLayoutError",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "2014-11-12T11:45:260Z",
				},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         time.Time{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseTime",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         testTime,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseDateTime",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testDateTime},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         testDateTime.Time(),
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 1},
				Key:          "value",
				DefaultValue: time.Time{},
			},
			output: outputStruct{
				Value:         time.Time{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := Date(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on Date(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.Value,
					gotValue,
				)
			}
			if fmt.Sprint(gotValid) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on Date(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.IsValid,
					gotValid,
				)
			}

			if fmt.Sprint(missingFieldsGot) != fmt.Sprint(tt.output.MissingFields) {
				t.Errorf(
					"expected missingFields on Date(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.MissingFields,
					missingFieldsGot,
				)
			}
		})
	}
}
