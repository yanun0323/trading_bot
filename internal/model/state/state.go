package state

type State uint8

const (
	Trading State = iota + 1
	Complete
)
