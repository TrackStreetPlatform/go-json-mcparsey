package Required

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue string
		}
		output struct {
			Value          string
			IsValid        bool
			RequiredFields []string
		}
	}{
		{
			name: "NoKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue string
			}{
				Origin: map[string]interface{}{
					"use_tosdfsdr": "1",
				},
				Key:          "use_tor",
				DefaultValue: "",
			},
			output: struct {
				Value          string
				IsValid        bool
				RequiredFields []string
			}{
				Value:   "",
				IsValid: false,
				RequiredFields: []string{
					"use_tor",
				}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requiredFieldsGot := []string{}
			valueGot, isValidGot := String(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on String(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on String(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on String(%v,%v,requiredFields,%v) = %v; got %v",
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

//func String(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue string) (value string, isValid bool) {
//	if maybeValueInField, ok := origin[key]; ok {
//		switch tempValueInField := maybeValueInField.(type) {
//		case string:
//			return tempValueInField, true
//		default:
//			AppendWhenNotNil(requiredFields, key)
//		}
//	} else {
//		AppendWhenNotNil(requiredFields, key)
//	}
//	return defaultValue, false
//}
