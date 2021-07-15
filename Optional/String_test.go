package Optional

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue string
		}
		output string
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue string
			}{
				Origin: map[string]interface{}{
					"value": "42",
				}, Key: "NonExisting", DefaultValue: ""},
			output: "",
		},
		{
			name: "CaseString",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue string
			}{
				Origin: map[string]interface{}{
					"value": "42",
				}, Key: "value", DefaultValue: ""},
			output: "42",
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue string
			}{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value", DefaultValue: ""},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected String(%v,%v,%v) = %v; got %v",
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
