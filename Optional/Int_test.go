package Optional

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue int
		}
		output int
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "NonExisting", DefaultValue: 0},
			output: 0,
		},
		{
			name: "CaseString",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": "42",
				}, Key: "value", DefaultValue: 0},
			output: 42,
		},
		{
			name: "CaseStringError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": "forty two",
				}, Key: "value", DefaultValue: 0},
			output: 0,
		},
		{
			name: "CaseInt",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "value", DefaultValue: 0},
			output: 42,
		},
		{
			name: "CaseInt32",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": int32(42),
				}, Key: "value", DefaultValue: 0},
			output: 42,
		},
		{
			name: "CaseInt64",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": int64(42),
				}, Key: "value", DefaultValue: 0},
			output: 42,
		},
		{
			name: "CaseFloat64",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": 42.,
				}, Key: "value", DefaultValue: 0},
			output: 42,
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue int
			}{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value", DefaultValue: 0},
			output: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected Int(%v,%v,%v) = %v; got %v",
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
