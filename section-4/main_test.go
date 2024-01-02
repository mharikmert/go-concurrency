package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		eatOrder = []Philosopher{}
		dine()
		if len(eatOrder) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(eatOrder))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		eatOrder = []Philosopher{}

		eatTime = e.delay
		thinkTime = e.delay

		dine()
		if len(eatOrder) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5 but got %d", e.name, len(eatOrder))
		}
	}
}
