package Optional

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"testing"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUUID(t *testing.T) {
	//testUUID := uuid.MustParse("abcdef01-0123-4567-89ab-0123456789ab")
	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue uuid.UUID
		}
		output uuid.UUID
	}{
		{
			name: "NonExistingKey",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": "1,2,3,TOR,luminati",
				}, Key: "NonExisting", DefaultValue: testUUID},
			output: testUUID,
		},
		{
			name: "StringCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": "00000000-0000-0000-0000-000000000000",
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},
		{
			name: "StringCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": "not a valid UUID",
				}, Key: "key", DefaultValue: testUUID},
			output: testUUID,
		},
		{
			name: "BinaryCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": primitive.Binary{Data: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},
		{
			name: "BinaryCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": primitive.Binary{Data: []byte{1, 1, 1, 1}},
				}, Key: "key", DefaultValue: testUUID},
			output: testUUID,
		},
		{
			name: "ByteArrayCase",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},
		{
			name: "ByteArrayCaseError",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": []byte{1, 1, 1, 1},
				}, Key: "key", DefaultValue: testUUID},
			output: testUUID,
		},
		{
			name: "CaseUUID",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": testUUID,
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},
		{
			name: "CaseUUIDMongo",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": uuid2.UUID{},
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},

		{
			name: "CaseUUIDMongo",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue uuid.UUID
			}{
				Origin: map[string]interface{}{
					"key": uuid2.UUID{},
				}, Key: "key", DefaultValue: uuid.New()},
			output: testUUID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UUID(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected UUID(%v,%v,%v) = %v; got %v",
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
