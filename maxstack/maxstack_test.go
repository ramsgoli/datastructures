package maxstack

import (
	"testing"
)

func TestPush(t *testing.T) {
	maxStack := MaxStack{}
	maxStack.Push(1)

	if len(maxStack.stack) != 1 {
		t.Errorf("Length of stack is not 1. Got %v instead\n", len(maxStack.stack))
	}

	if len(maxStack.maxes) != 1 {
		t.Errorf("Length of maxes is not 1. Got %v instead\n", len(maxStack.maxes))
	}
}

func TestPop(t *testing.T) {
	maxStack := MaxStack{}
	maxStack.Push(1)
	maxStack.Push(2)

	popped, err := maxStack.Pop()

	if err != nil {
		t.Error("Got an error popping")
	}

	if popped != 2 {
		t.Errorf("Expected %v, got %v\n", 2, popped)
	}

	if len(maxStack.stack) != 1 {
		t.Errorf("Expected length after pop is %v, got %v\n", 1, len(maxStack.stack))
	}
}

func TestMax(t *testing.T) {
	maxStack := MaxStack{}
	maxStack.Push(1)
	maxStack.Push(2)

	if maxStack.Max() != 2 {
		t.Error("Invalid max")
	}
}
