package gofp

type (
	Function[I, O any]    func(I) O
	Predicate[T any]      func(T) bool
	Accumulator[R, I any] func(R, I) R
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

func Reduce[T1, T2 any](i []T1, seed T2, acc Accumulator[T2, T1]) T2 {
	a := seed
	for _, e := range i {
		a = acc(a, e)
	}
	return a
}
