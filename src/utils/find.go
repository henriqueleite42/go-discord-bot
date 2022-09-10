package utils

import "errors"

func Find[T any](arr []T, compare func(T) bool) (T, error) {
	for _, value := range arr {
		if compare(value) {
			return value, nil
		}
	}

	return arr[0], errors.New("not found")
}
