package Required

import (
	"fmt"
	"testing"
)

type inputStruct struct {
	Origin         map[string]interface{}
	Key            string
	RequiredFields *[]string
	DefaultValue   []string
}

type outputStruct struct {
	Value   []string
	IsValid bool
}

type testArrayString struct {
	name   string
	input  inputStruct
	output outputStruct
}

func TestArrayString(t *testing.T) {
	tests := []testArrayString{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:         map[string]interface{}{"value": "1"},
				Key:            "NonExisting",
				RequiredFields: &[]string{},
				DefaultValue:   []string{},
			},
			output: outputStruct{
				Value:   []string{},
				IsValid: false,
			},
		},
		{
			name: "NonExistentKeysAppend",
			input: inputStruct{
				Origin:         map[string]interface{}{"value": "1"},
				Key:            "NonExisting",
				RequiredFields: &[]string{"not nil"},
				DefaultValue:   []string{},
			},
			output: outputStruct{
				Value:   []string{},
				IsValid: false,
			},
		},
		//{
		//	name: "CaseString",
		//	input: inputStruct{
		//		Origin: map[string]interface{}{
		//			"value": "1,2",
		//		},
		//		Key: "value",
		//		RequiredFields: &[]string{"value"} ,
		//		DefaultValue: []string{},
		//	},
		//	output: outputStruct{
		//		Value: []string{"1","2"},
		//		IsValid: true,
		//	},
		//},
		//{
		//	name: "UnsupportedType",
		//	input: inputStruct{
		//		Origin: map[string]interface{}{"value": 42,},
		//		Key: "value",
		//		DefaultValue: []string{},
		//	},
		//	output: outputStruct{
		//		Value: []string{},
		//		IsValid: false,
		//	},
		//},
		//{
		//	name: "InterfaceArrayProper",
		//	input: inputStruct{
		//		Origin: map[string]interface{}{
		//			"value": []interface{}{"disused", "sdfsd"},
		//		}, Key: "value", DefaultValue: []string{}},
		//	output: outputStruct{
		//		Value: []string{},
		//		IsValid: false,
		//	},
		//},
		//{
		//	name: "InterfaceArrayProperOneItemWrongType",
		//	input: inputStruct{
		//		Origin: map[string]interface{}{
		//			"value": []interface{}{"disused", "sdfsd", 42},
		//		}, Key: "value", DefaultValue: []string{}},
		//	output: outputStruct{
		//		Value: []string{},
		//		IsValid: false,
		//	},
		//},
		//{
		//	name: "StringArrayProper",
		//	input: inputStruct{
		//		Origin: map[string]interface{}{
		//			"value": []string{"disused", "sdfsd"},
		//		}, Key: "value", DefaultValue: []string{}},
		//	output: outputStruct{
		//		Value: []string{},
		//		IsValid: false,
		//	},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotValid := ArrayString(tt.input.Origin, tt.input.Key, tt.input.RequiredFields, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) || tt.output.IsValid != gotValid {
				t.Errorf(
					"expected ArrayString(%v,%v,%v,%v) = %v,%v; got %v, %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.RequiredFields,
					tt.input.DefaultValue,
					tt.output.Value,
					tt.output.IsValid,
					gotValue,
					gotValid,
				)
			}
		})
	}

}
