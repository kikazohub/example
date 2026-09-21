package reverse_test

import (
	"testing"

	"golang.org/x/example/hello/reverse"
)

func TestInt(t *testing.T) {
	for _, c := range []struct {
		in, want int
	}{
		{0, 0},
		{1, 1},
		{12, 21},
		{123, 321},
		{100, 1},
		{-123, -321},
		{-1, -1},
	} {
		got := reverse.Int(c.in)
		if got != c.want {
			t.Errorf("Int(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
