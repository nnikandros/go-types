package types

// wrapper struct
type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {

	return Result[T]{Value: value}

}

func Error[T any](err error) Result[T] {

	return Result[T]{Err: err}

}

func Wrap[T any](v T, err error) Result[T] {
	return Result[T]{Value: v, Err: err}

}

func UnWrap[T any](v Result[T]) (T, error) {
	return v.Value, v.Err

}
