package action

type Action uint8

const (
	None Action = iota + 1
	Buy
	Sell
)
