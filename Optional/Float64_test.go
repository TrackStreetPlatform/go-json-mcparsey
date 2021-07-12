package Optional

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue float64
		}
		output float64
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue float64
			}{
				Origin: map[string]interface{}{
					"value": 42.0,
				}, Key: "NonExisting", DefaultValue: 0.0},
			output: 0.0,
		},
		{
			name: "CaseInt",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue float64
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "value", DefaultValue: 0.},
			output: 42.,
		},
		{
			name: "CaseFloat",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue float64
			}{
				Origin: map[string]interface{}{
					"value": 42.,
				}, Key: "value", DefaultValue: 0.},
			output: 42.,
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue float64
			}{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value", DefaultValue: 0.},
			output: 0.,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected Float64(%v,%v,%v) = %v; got %v",
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
