package garrays

func Search[T any](array []T, predicate func(element T, index int) bool) []T {
	filteredData := make([]T, 0)
	for i, element := range array {
		if predicate(element, i) {
			filteredData = append(filteredData, element)
		}
	}
	return filteredData
}
