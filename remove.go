package garrays

func Remove[T any](array []T, index int) []T {
	for i := range array {
		if i == index {
			return append(array[:i], array[i+1:]...)
		}
	}
	return array
}

func Filter[T any](array []T, predicate func(element T, index int) bool) []T {
	filteredData := make([]T, 0)
	for i, element := range array {
		if !predicate(element, i) {
			filteredData = append(filteredData, element)
		}
	}
	return filteredData
}

func Splice[T any](array []T, start int, deleteCount int) []T {
	end := deleteCount + start
	for i := range array {
		if start <= i && i < end {
			return append(array[:i], array[i+deleteCount:]...)
		}
	}
	return array
}
