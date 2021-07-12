package Required

func AppendWhenNotNil(requiredFields *[]string, key string) {
	if requiredFields != nil {
		*requiredFields = append(*requiredFields, key)
	}
}
