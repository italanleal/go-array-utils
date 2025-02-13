package utils

func Map[T any, R any](arr []T, fn func(T) R) []R {
	var result []R
	for _, value := range arr {
		result = append(result, fn(value))
	}
	return result
}

func Reduce[T any](arr []T, fn func(acc T, cur T) T, start T) T {
	var result T = start
	for _, value := range arr {
		result = fn(result, value)
	}
	return result
}

func ForEach[T any](arr []T, fn func(T)) {
	for _, value := range arr {
		fn(value)
	}
}

func Reverse[T any](arr []T) []T {
	var result []T

	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}

	return result
}
