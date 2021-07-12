# go-json-mcparsey

go-json-mcparsey is a very forgiving json helper parser for golang, the default one is very strict when parsing to a struct and would rightfully break when the json is not using the exact correct type, this tries as much as possible to parse to the correct type, going out of it's way to accept a json that was not parse 100% correctly
