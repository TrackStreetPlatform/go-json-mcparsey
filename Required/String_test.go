package Required

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue string
	}
	type outputStruct struct {
		Value         string
		IsValid       bool
		MissingFields []string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "NonExistingKey",
				DefaultValue: "",
			},
			output: outputStruct{
				Value:         "",
				IsValid:       false,
				MissingFields: []string{"NonExistingKey"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "value",
				DefaultValue: "",
			},
			output: outputStruct{
				Value:         "1",
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 1},
				Key:          "value",
				DefaultValue: "",
			},
			output: outputStruct{
				Value:         "",
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			valueGot, isValidGot := String(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on String(%v,%v,missingFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on String(%v,%v,missingFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(missingFieldsGot) != fmt.Sprint(tt.output.MissingFields) {
				t.Errorf(
					"expected missingFields on String(%v,%v,missingFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.MissingFields,
					missingFieldsGot,
				)
			}

		})
	}
}

//func String(origin map[string]interface{}, key string, missingFields *[]string, defaultValue string) (value string, isValid bool) {
//	if maybeValueInField, ok := origin[key]; ok {
//		switch tempValueInField := maybeValueInField.(type) {
//		case string:
//			return tempValueInField, true
//		default:
//			AppendNotNil(missingFields, key)
//		}
//	} else {
//		AppendNotNil(missingFields, key)
//	}
//	return defaultValue, false
//}
