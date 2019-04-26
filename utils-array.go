package gosf

// ArrayContainsString checks to see if a value exists in the array
func ArrayContainsString(array []string, value string) bool {
	hasValue := false
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			hasValue = true
		}
	}
	return hasValue
}
