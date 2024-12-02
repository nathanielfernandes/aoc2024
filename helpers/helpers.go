package helpers

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func Abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
