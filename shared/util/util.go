package util

import (
	"golang.org/x/exp/constraints"
)

func Increment[T constraints.Integer]() <-chan T {
	var i T = 0
	result := make(chan T, 64)
	go func() {
		for {
			result <- i
			i++
		}
	}()
	return result
}

func EventHandle[T any](ch <-chan T, action func(T)) {
	for {
		select {
		case evt := <-ch:
			action(evt)
		default:
			return
		}
	}
}
