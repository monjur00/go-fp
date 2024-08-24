package gofp

type Monad[T any] interface {
	Bind(BindFn[T]) Monad[T]
}

type BindFn[T any] func(T) Monad[T]

type Try[T any] struct {
	Result T
	Err    error
}

func (t Try[T]) Bind(b BindFn[T]) Monad[T] {
	if t.Err != nil {
		return &Try[T]{t.Result, t.Err}
	}

	return b(t.Result)
}

func NewTry[T any](t T, err error) Try[T] {
	return Try[T]{t, err}
}

// TODO: Either[R, L any] for more generic options
type Either[R any] struct {
	Right R
	Left  error
}

func NewEither[R any](r R, l error) Either[R] {
	return Either[R]{r, l}
}

func (e Either[R]) Bind(b BindFn[R]) Monad[R] {
	if e.Left != nil {
		return &Either[R]{e.Right, e.Left}
	}

	return b(e.Right)
}
