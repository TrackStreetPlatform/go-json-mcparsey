package Optional

import "testing"

func TestString(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output string
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "42",
				}, Key: "NonExisting", DefaultValue: ""},
			output: "",
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "42",
				}, Key: "value", DefaultValue: ""},
			output: "42",
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
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
