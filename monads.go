package gofp

type Monad[T any] interface {
	Bind(BindFn[T]) Monad[T]
}

type BindFn[T any] func(T) Monad[T]

type Try[T any] struct {
	Result T
	Err    error
}

func (t *Try[T]) Bind(b BindFn[T]) Monad[T] {
	if t.Err != nil {
		return &Try[T]{t.Result, t.Err}
	}

	return b(t.Result)
}

func NewTry[T any](t T, err error) Try[T] {
	return Try[T]{t, err}
}
