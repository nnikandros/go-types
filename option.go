package types

type Option[T any] struct {
	Value T
	IsSet bool
}
