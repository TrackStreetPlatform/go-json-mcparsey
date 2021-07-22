package Required

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue uuid.UUID
	}

	type outputStruct struct {
		Value          uuid.UUID
		IsValid        bool
		RequiredFields []string
	}

	defaultValue, _ := uuid.Parse("01234567-8910-1112-1314-151617181920") // Random bytes
	correctValue, _ := uuid.Parse("636f7272-6563-7456-616c-756555554944") // Generated from bytes "correctValueUUID"

	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NoKey",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "636f7272-6563-7456-616c-756555554944",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "invalid string",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidBinary",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": primitive.Binary{Data: []byte("correctValueUUID")},
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidBinary",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": primitive.Binary{Data: []byte("bad uuid")},
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidBytes",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []byte("correctValueUUID"),
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidBytes",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []byte("bad uuid"),
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "UUID",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": correctValue,
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "UUID2",  // UUID2 is more strict than UUID, so there's no case for an invalid UUID2
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": uuid2.UUID([16]byte{
						0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x56,
						0x61, 0x6c, 0x75, 0x65, 0x55, 0x55, 0x49, 0x44,
					}),
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": false,
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := UUID(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on UUID(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on UUID(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on UUID(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.RequiredFields,
					requiredFieldsGot,
				)
			}

		})
	}
}
