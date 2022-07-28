package Required

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUUID(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue uuid.UUID
	}
	type outputStruct struct {
		Value         uuid.UUID
		IsValid       bool
		MissingFields []string
	}
	testUUID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testUUID},
				Key:          "NonExisting",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         uuid.UUID{},
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": "11111111-1111-1111-1111-111111111111",
				},
				Key:          "value",
				DefaultValue: testUUID,
			},
			output: outputStruct{
				Value:         testUUID,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseStringErr",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "not valid"},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         uuid.UUID{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseBinary",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": primitive.Binary{
						Data: []byte{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
					},
				},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         testUUID,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseBinaryError",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": primitive.Binary{
						Data: []byte{1, 1, 1, 1},
					},
				},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         uuid.UUID{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseByteArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []byte{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
				},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         testUUID,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseByteArrayError",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []byte{1, 1, 1, 1}},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         uuid.UUID{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseUUID",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": testUUID},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         testUUID,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{"1"}},
				Key:          "value",
				DefaultValue: uuid.UUID{},
			},
			output: outputStruct{
				Value:         uuid.UUID{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := UUID(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on UUID(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.Value,
					gotValue,
				)
			}
			if fmt.Sprint(gotValid) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on UUID(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.IsValid,
					gotValid,
				)
			}
			if fmt.Sprint(missingFieldsGot) != fmt.Sprint(tt.output.MissingFields) {
				t.Errorf(
					"expected missingFields on UUID(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.MissingFields,
					missingFieldsGot,
				)
			}
		})
	}
}
