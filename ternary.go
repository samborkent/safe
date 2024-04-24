package safe

func Ternary[T any](condition *bool, incase T, otherwise T) T {
	if *condition {
		return incase
	} else {
		return otherwise
	}
}
