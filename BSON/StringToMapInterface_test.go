package JSON

import (
	"errors"
	"fmt"
	"testing"
)

func TestStringToMapInterface(t *testing.T) {
	type inputStruct struct {
		MessageBody []byte
		TypeStr     string
	}
	type outputStruct struct {
		Value map[string]interface{}
		Error error
	}
	//bsonStruct := bson.M{"test": bson.M{"value": 42}}
	//testJSON, _ := bson.Marshal(bsonStruct)

	//rec := make([]byte, len(testJSON)+6)
	//binary.LittleEndian.PutUint32(rec, uint32(len(rec)))
	//copy(rec[4:], testJSON)

	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "ErrorParsing",
			input: inputStruct{
				MessageBody: []byte("invalid"),
				TypeStr:     "test",
			},
			output: outputStruct{
				Value: nil,
				Error: errors.New("invalid document length"),
			},
		},
		//{
		//	name: "ProperParsing",
		//	input: inputStruct{
		//		MessageBody:	testJSON,
		//		TypeStr:        "test",
		//	},
		//	output: outputStruct{
		//		Value:         	map[string]interface{}{
		//			"test": 	map[string]interface{}{"value": 42},
		//		},
		//		Error: 			nil,
		//	},
		//},
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
