package slices

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, filter func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
