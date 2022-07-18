package Path

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `
	{
		"number": 21,
		"text": "this is a string",
		"boolean": true,
		"obj": {
			"text": "text",
			"10": 10,
			"string_list": ["one", "two", "three"],
			"obj_list": [
				{"name": "one"},
				{"name": "two", "innerlist": [{"n":"aa1"}, {"n":"aa2"}, {"n":"aa3"}]},
				{"name": "three"},
				{"name": "four", "innerlist": [{"n":"ii1"}, {"n":"ii2"}, {"n":"ii3"}, {"n":"ii4"}, {"n":"ii5"}]}
			]
		}
	}
`

func TestTraverse(t *testing.T) {
	var tree interface{}
	err := json.Unmarshal([]byte(jsonStr), &tree)
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}

	type testCase struct {
		name     string
		path     string
		expected interface{}
	}

	testCases := []testCase{
		{
			name:     "find field in level 1",
			path:     "number",
			expected: 21,
		},
		{
			name:     "nonexistent field",
			path:     "nonexistent",
			expected: nil,
		},
		{
			name:     "find field in level 2",
			path:     "obj > text",
			expected: "text",
		},
		{
			name:     "find fields in list",
			path:     "obj > obj_list > name",
			expected: []string{"one", "two", "three", "four"},
		},
		{
			name:     "find obj in list with index [2]",
			path:     "obj > obj_list[2]",
			expected: map[string]string{"name": "three"},
		},
		{
			name:     "find field in list with index [2]",
			path:     "obj > obj_list[2] > name",
			expected: "three",
		},
		{
			name:     "return nil when index is out of bounds [-1]",
			path:     "obj > obj_list[-1]",
			expected: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Traverse(tree, tt.path)
			if fmt.Sprint(got) != fmt.Sprint(tt.expected) {
				t.Errorf("using path: (%s) expected result=%v; got=%v", tt.path, tt.expected, got)
			}
		})
	}
}

func Test_querySteps(t *testing.T) {
	type testCase struct {
		name     string
		query    string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "chain of fields",
			query:    "field > field2 > field3",
			expected: []string{"field", "field2", "field3"},
		},
		{
			name:     "index syntax",
			query:    "list[2] > name",
			expected: []string{"list", "[2]", "name"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			steps := querySteps(tt.query)
			if fmt.Sprint(steps) != fmt.Sprint(tt.expected) {
				t.Errorf("expected result=%v; got=%v", tt.expected, steps)
			}
		})
	}
}
