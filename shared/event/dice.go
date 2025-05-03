package event

type DiceEvent[T any] struct {
	Id uint64
	DiceEventType
	Result T
}

func (d DiceEvent[T]) RequestId() uint64 {
	return d.Id
}

type DiceEventType int

const (
	DiceEventTypeResult DiceEventType = iota
)
