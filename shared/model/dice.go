package model

import "math/rand"

type DiceModel[T any] interface {
	Roll()
	Result() T
	ResultScreen() int
	RollAble() bool
	GetStatus() DiceStatus
	SetStatus(DiceStatus)
}

type DiceStatus int

const (
	DiceStatusReady DiceStatus = iota
	DiceStatusRolling
	DiceStatusRoll
	DiceStatusResult
)

type BaseDice[T any] struct {
	status DiceStatus
	Value  int
	Values []T
}

func (b *BaseDice[T]) Roll() {
	b.Value = rand.Intn(6)
}

func (b *BaseDice[T]) Result() T {
	return b.Values[b.Value]
}

func (b *BaseDice[T]) ResultScreen() int {
	return b.Value
}

func (b *BaseDice[T]) RollAble() bool {
	return b.GetStatus().RollAble()
}

func (d DiceStatus) RollAble() bool {
	return d != DiceStatusRoll
}

func (b *BaseDice[T]) GetStatus() DiceStatus {
	return b.status
}

func (b *BaseDice[T]) SetStatus(status DiceStatus) {
	b.status = status
}
