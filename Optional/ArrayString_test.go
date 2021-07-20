package Optional

import (
	"fmt"
	"testing"
)

func TestArrayString(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue []string
		}
		output []string
	}{

		{
			name: "NonExistentKeys",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": "1",
				}, Key: "NonExisting", DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "CaseString",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": "1,2",
				}, Key: "value", DefaultValue: []string{}},
			output: []string{"1", "2"},
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "value", DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "InterfaceArrayProper",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{"disused", "sdfsd"},
				}, Key: "value", DefaultValue: []string{}},
			output: []string{"disused", "sdfsd"},
		},
		{
			name: "InterfaceArrayProperOneItemWrongType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{"disused", "sdfsd", 42},
				}, Key: "value", DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "StringArrayProper",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"value": []string{"disused", "sdfsd"},
				}, Key: "value", DefaultValue: []string{}},
			output: []string{"disused", "sdfsd"},
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
