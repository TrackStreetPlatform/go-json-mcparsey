package Required

func AppendNotNil(missingFields *[]string, key string) {
	if missingFields != nil {
		*missingFields = append(*missingFields, key)
	}
}
