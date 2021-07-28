package Optional

import (
	"fmt"
	"testing"
)

func TestInterfaceArray(t *testing.T) {
	type inputStruct struct {
		Origin map[string]interface{}
		Key    string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output []map[string]interface{}
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{map[string]interface{}{"test": 42}},
				},
				Key: "NonExisting",
			},
			output: []map[string]interface{}{},
		},
		{
			name: "ListInterface",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{map[string]interface{}{"test": 42}},
				},
				Key: "value",
			},
			output: []map[string]interface{}{{"test": 42}},
		},
		{
			name: "ListInterfaceMultipleEntries",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{
						map[string]interface{}{"test1": 42},
						map[string]interface{}{"test2": "42"}},
				},
				Key: "value",
			},
			output: []map[string]interface{}{{"test1": 42}, {"test2": "42"}},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []string{},
				},
				Key: "value",
			},
			output: []map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InterfaceArray(tt.input.Origin, tt.input.Key)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected InterfaceArray(%v,%v) = %v got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.output,
					got,
				)
			}
		})
	}
}
