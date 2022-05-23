package Optional

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func InterfaceArray(origin map[string]interface{}, path string) []map[string]interface{} {
	items := make([]map[string]interface{}, 0)
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case []interface{}:
			for _, maybeItem := range tempValueInField {
				switch tempValueInField := maybeItem.(type) {
				case map[string]interface{}:
					items = append(items, tempValueInField)
				default:
					break
				}

			}
		default:
			break
		}
	}
	return items
}
