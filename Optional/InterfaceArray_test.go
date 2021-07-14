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
		//[]map[string]interface{}{{"proxies": "1,2,3"}, {"attributes": "Cache"}}
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "NonExisting"},
			output: make([]map[string]interface{}, 0),
		},
		{
			name: "CommonCase",
			input: struct {
				Origin map[string]interface{}
				Key    string
			}{
				Origin: map[string]interface{}{
					"value": []interface{}{"testing1", "testing2"},
				}, Key: "value"},
			output: []map[string]interface{}{},
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
			//fmt.Printf("Got:\t %v, %v\n", got, reflect.TypeOf(got))
			//fmt.Printf("Output:\t %v, %v\n", tt.output, reflect.TypeOf(tt.output))
			//fmt.Println(reflect.DeepEqual(got, tt.output))
			if reflect.DeepEqual(got, tt.output) {
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
