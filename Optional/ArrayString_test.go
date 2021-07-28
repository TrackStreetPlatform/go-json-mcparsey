package Optional

import (
	"fmt"
	"testing"
)

func TestArrayString(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue []string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output []string
	}{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "NonExisting",
				DefaultValue: []string{},
			},
			output: []string{},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1,2"},
				Key:          "value",
				DefaultValue: []string{}},
			output: []string{"1", "2"},
		},
		{
			name: "CaseInterfaceArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{"test", "42"}},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{"test", "42"},
		},
		{
			name: "CaseInterfaceArrayError",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{"test", 42}},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{},
		},
		{
			name: "CaseStringArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []string{"test", "42"},
				},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{"test", "42"},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42},
				Key:          "value",
				DefaultValue: []string{}},
			output: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayString(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected ArrayString(%v,%v,%v) = %v; got %v",
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
