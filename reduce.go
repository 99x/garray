package garray

func Reduce[T, M any](s []T, f func(M, T) M, initial M) M {
	current := initial
	for _, v := range s {
		current = f(current, v)
	}
	return current
}
