// Copyright 2026 Enrique Vazquez. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2026 Enrique Vazquez. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse_test

import (
	"testing"

	"github.com/kikazohub/example/hello/reverse"
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
