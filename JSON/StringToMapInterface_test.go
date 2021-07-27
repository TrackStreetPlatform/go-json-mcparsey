package JSON

import (
	"errors"
	"fmt"
	"testing"
)

func TestStringToMapInterface(t *testing.T) {
	type inputStruct struct {
		MessageBody string
		TypeStr     string
	}
	type outputStruct struct {
		Value map[string]interface{}
		Error error
	}
	testJSON := "{\"test\": {\n  \"value\": 42\n}}"
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "ErrorParsing",
			input: inputStruct{
				MessageBody: "invalid",
				TypeStr:     "test",
			},
			output: outputStruct{
				Value: nil,
				Error: errors.New("invalid character 'i' looking for beginning of value"),
			},
		},
		{
			name: "ProperParsing",
			input: inputStruct{
				MessageBody: testJSON,
				TypeStr:     "test",
			},
			output: outputStruct{
				Value: map[string]interface{}{
					"test": map[string]interface{}{"value": 42},
				},
				Error: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotError := StringToMapInterface(tt.input.MessageBody, tt.input.TypeStr)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on StringToMapInterface(%v,%v) = %v; got %v",
					tt.input.MessageBody,
					tt.input.TypeStr,
					tt.output.Value,
					gotValue,
				)
			}
			if fmt.Sprint(gotError) != fmt.Sprint(tt.output.Error) {
				t.Errorf(
					"expected error on StringToMapInterface(%v,%v) = %v; got %v",
					tt.input.MessageBody,
					tt.input.TypeStr,
					tt.output.Error,
					gotError,
				)
			}
		})
	}
}
