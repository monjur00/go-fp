package gofp

type (
	Function[I, O any] func(I) O
	Predicate[T any]   func(T) bool
)

func Map[I, O any](i []I, f Function[I, O]) []O {
	var o []O
	for _, e := range i {
		o = append(o, f(e))
	}
	return o
}

func Filter[I any](i []I, p Predicate[I]) []I {
	var o []I
	for _, e := range i {
		if p(e) {
			o = append(o, e)
		}
	}
	return o
}
