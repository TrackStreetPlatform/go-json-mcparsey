package Path

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	keyIndexPattern   = regexp.MustCompile(`(\w+)(\[\d+\])`)
	indexTokenPattern = regexp.MustCompile(`\[(\d+)\]`)
)

const fieldSeparator = ">"

// Traverse the provided interface with the given path. It returns the field at the end of the path.
// In order to make a traversable path, every node name must be separated by the '>' symbol. ex: field1 > field2
// To visit a specific element in a list, use the [number] notation. ex: field_name[2]
func Traverse(tree interface{}, path string) (interface{}, error) {
	var err error
	var iterated bool
	return traverse(tree, &err, &iterated, querySteps(path))
}

func traverse(node interface{}, err *error, iterated *bool, path []string) (interface{}, error) {
	if node == nil || len(path) == 0 || *err != nil {
		return node, *err
	}
	headKey := path[0]

	switch cnode := node.(type) {
	case map[string]interface{}:
		newNode, ok := cnode[headKey]
		if !ok {
			newNode = nil
		}
		return traverse(newNode, err, iterated, path[1:])
	case []interface{}:
		if matches := indexTokenPattern.FindStringSubmatch(headKey); len(matches) > 1 {
			idx, _ := strconv.Atoi(matches[1])
			var tmpNode interface{}
			if idx < 0 || idx >= len(cnode) {
				tmpE := fmt.Errorf("index %d out of bounds", idx)
				err = &tmpE
			} else {
				tmpNode = cnode[idx]
			}
			return traverse(tmpNode, err, iterated, path[1:])
		} else {
			*iterated = true
			var list []interface{}
			for _, elem := range cnode {
				if nt, ok := elem.(map[string]interface{}); ok {
					n, err1 := traverse(nt, err, iterated, path)
					if err1 != nil {
						n = nil
					}
					if nl, isList := n.([]interface{}); isList && false { // disable
						list = append(list, nl...)
					}
					list = append(list, n)
				}
			}
			if len(list) == 0 {
				return nil, *err
			}
			return list, *err
		}
	default:
		tmpE := fmt.Errorf("unknown type")
		err = &tmpE
	}

	return nil, *err
}

func querySteps(query string) (steps []string) {
	if strings.Contains(query, "|") {
		return
	}
	for _, step := range strings.Split(query, fieldSeparator) {
		if matches := keyIndexPattern.FindStringSubmatch(step); len(matches) > 2 {
			steps = append(steps, strings.TrimSpace(matches[1]))
			steps = append(steps, strings.TrimSpace(matches[2]))
			continue
		}
		steps = append(steps, strings.TrimSpace(step))
	}
	return
}
