package event

type Event interface {
	RequestId() uint64
}
