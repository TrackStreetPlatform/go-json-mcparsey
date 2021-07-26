package Optional

import "testing"

func TestBool(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue bool
	}
	tests := []struct {
		name   string
		input  inputStruct
		output bool
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": true,
				}, Key: "NonExisting", DefaultValue: false},
			output: false,
		},
		{
			name: "CaseStringTrue",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "true",
				}, Key: "value", DefaultValue: false},
			output: true,
		},
		{
			name: "CaseString1",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "1",
				}, Key: "value", DefaultValue: false},
			output: true,
		},
		{
			name: "CaseInteger",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": 1,
				}, Key: "value", DefaultValue: false},
			output: true,
		},
		{
			name: "CaseFloat",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": 1.0,
				}, Key: "value", DefaultValue: false},
			output: true,
		},
		{
			name: "CaseBool",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": true,
				}, Key: "value", DefaultValue: false},
			output: true,
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value", DefaultValue: false},
			output: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bool(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected Bool(%v,%v,%v) = %v; got %v",
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
