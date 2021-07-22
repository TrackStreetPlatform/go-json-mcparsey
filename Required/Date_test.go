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
		Value          time.Time
		IsValid        bool
		RequiredFields []string
	}

	defaultValue := time.Date(2002, 3, 14, 3, 14, 15, 92, time.UTC)
	correctValue := time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)

	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NoKey",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ISOTimeString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "2001-02-03T04:05:06Z",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "NonstandardTimeString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "2001-02-03 04:05:06",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "This isn't valid",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "MongoTime",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": primitive.NewDateTimeFromTime(correctValue),
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []int{1, 2, 3},
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := Date(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if !valueGot.Equal(tt.output.Value) { // Mongo changes timezones, so we need to compare Dates exactly
				t.Errorf(
					"expected value on Date(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on Date(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on Date(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.RequiredFields,
					requiredFieldsGot,
				)
			}

		})
	}
}
