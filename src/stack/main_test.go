package stack_test

import (
	"testing"
	"github.com/sreesindhusruthiyadavalli/go-alds/src/stack"
)

func benchmarkCalculate(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack.CheckStack()
	}
}

func BenchmkdimarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }