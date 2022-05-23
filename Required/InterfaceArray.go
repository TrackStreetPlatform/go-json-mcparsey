package Required

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func InterfaceArray(origin map[string]interface{}, path string, missingFields *[]string) (value []map[string]interface{}, isValid bool) {
	items := make([]map[string]interface{}, 0)
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case []interface{}:
			for _, maybeItem := range tempValueInField {
				switch tempValueInField := maybeItem.(type) {
				case map[string]interface{}:
					items = append(items, tempValueInField)
					isValid = true
				default:
					break
				}
			}
		default:
			AppendNotNil(missingFields, path)
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return items, isValid
}
