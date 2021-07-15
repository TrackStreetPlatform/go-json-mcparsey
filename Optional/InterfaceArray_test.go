package Optional

import (
	"reflect"
	"testing"
)

func TestInterfaceArray(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin map[string]interface{}
			Key    string
		}
		output []map[string]interface{}
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{map[string]interface{}{"test": 42}},
				}, Key: "NonExisting"},
			output: []map[string]interface{}{},
		},
		{
			name: "CommonCase",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{map[string]interface{}{"test": 42}},
				}, Key: "value"},
			output: []map[string]interface{}{{"test": 42}},
		},
		{
			name: "CommonCaseMultipleEntries",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{
						map[string]interface{}{"test1": 42},
						map[string]interface{}{"test2": "42"}},
				}, Key: "value"},
			output: []map[string]interface{}{{"test1": 42}, {"test2": "42"}},
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": []string{},
				}, Key: "value"},
			output: []map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InterfaceArray(tt.input.Origin, tt.input.Key)
			if !reflect.DeepEqual(got, tt.output) {
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
