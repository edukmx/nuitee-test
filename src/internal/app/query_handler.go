package app

type Query interface{}

type QueryHandler[I Query, O any] interface {
	Ask(query I) (*O, error)
}
