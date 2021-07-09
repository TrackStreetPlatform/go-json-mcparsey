package go_json_mcparsey

func AppendWhenNotNil(requiredFields *[]string, key string) {
	if requiredFields != nil {
		*requiredFields = append(*requiredFields, key)
	}
}
