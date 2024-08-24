package gofp

type (
	Monad[T any] interface {
		Bind(BindFn[T]) Monad[T]
	}

	BindFn[T any] func(T) Monad[T]

	Try[T any] struct {
		Result T
		Err    error
	}

	// TODO: Either[R, L any] for more generic options
	Either[R any] struct {
		Right R
		Left  error
	}

	Maybe[T comparable] struct {
		Value T
	}
)

func (t Try[T]) Bind(b BindFn[T]) Monad[T] {
	if t.Err != nil {
		return &Try[T]{t.Result, t.Err}
	}

	return b(t.Result)
}

func NewTry[T any](t T, err error) Try[T] {
	return Try[T]{t, err}
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

func NewMaybe[T comparable](t T) Maybe[T] {
	return Maybe[T]{t}
}

func (m Maybe[T]) Bind(b BindFn[T]) Monad[T] {
	var t T
	if m.Value == t {
		return &Maybe[T]{m.Value}
	}

	return b(m.Value)
}
