package sample

import (
    "testing"
)

func TestAdd(t *testing.T) {
    got := add(5, 3)
    if got != 8 {
        t.Errorf("add(5, 3) = %d; want 8", got)
    }
}

func TestSubtract(t *testing.T) {
    got := subtract(5, 3)
    if got != 2 {
        t.Errorf("subtract(5, 3) = %d; want 2", got)
    }
}
