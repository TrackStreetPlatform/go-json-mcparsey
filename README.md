# go-json-mcparsey

go-json-mcparsey is a very forgiving json helper parser for golang, the default one is very strict when parsing to a struct and would rightfully break when the json is not using the exact correct type, this tries as much as possible to parse to the correct type, going out of it's way to accept a json that was not stringfied using the exact type

## Installation

    go get github.com/TrackStreetPlatform/go-json-mcparsey

## Path Traversal

It is possible to get fields from different node levels by providing the keys leading to the desired node.
Similar to following breadcrumbs, each level must be separated by `>` fd.
If the encountered node is an Array, it is possible to select the element we want to take by specifying its index `[0]`.

Example JSON:

    {
        "menu": {
            "name": "something",
            "items": [
                {
                    "key": "one",
                    "value": 1
                },
                {
                    "key": "two",
                    "value": 2
                }
            ]
        },
        "date": "2023-06-21"
    }

If we decode the last json and dump it into a variable called decodedJson, we're able to read any field
from it without needing to check first if the key exists or creating temporal variables just to reach a lower level node.

    sm := Optional.String(decodedJson, "menu > name", "")                 // something
    one := Optional.Int(decodedJson, "menu > items[0] > key", 0)          // 1
    m := Optional.MapStringInterface(decodedJson, "menu > items[0]", nil) // map[string]interface{}

