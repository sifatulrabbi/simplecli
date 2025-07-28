package utils

func Ternary[T any](cond bool, ifTrue, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func TernaryFnCall[T any](cond bool, ifTrue, ifFalse func() T) T {
	if cond {
		ifTrue()
	}
	return ifFalse()
}
