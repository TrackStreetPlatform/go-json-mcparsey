package Optional

import (
	"fmt"
	"testing"
)

func TestDispatchPluginsFromQueue(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue []string
		}
		output []string
	}{

		{
			name: "ValidProxies",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": "1,2,3,TOR,luminati",
				}, Key: "proxies", DefaultValue: []string{}},
			output: []string{"1", "2", "3", "TOR", "luminati"},
		},
		{
			name: "KeyNotExists",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": "1,2,3,TOR,luminati",
				}, Key: "key_not_exists", DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "NotStringValue",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": 3443,
				}, Key: "proxies", DefaultValue: []string{}},
			output: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayString(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected ArrayString(%v,%v,%v) = %v; got %v",
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
