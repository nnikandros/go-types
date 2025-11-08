package types

type Option[T any] struct {
	Value T
	IsValid bool
}
