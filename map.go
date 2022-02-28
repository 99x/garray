package garrays

func Map[T any, E any](array []T, fn func(element T, index int) E) *[]E {
	returnArray := make([]E, 0)
	for i, element := range array {
		returnArray = append(returnArray, fn(element, i))
	}

	return &returnArray
}
