package garrays

func Find[T any](array []T, predicate func(element T) bool) *T {
	for _, element := range array {
		if predicate(element) {
			return &element
		}
	}
	return nil
}

func FindIndex[T any](array []T, predicate func(element T) bool) int {
	for i, element := range array {
		if predicate(element) {
			return i
		}
	}
	return -1
}
