package order

type Type uint8

const (
	Buy Type = iota + 1
	Sell
)
