// Copyright 2026 Enrique Vazquez. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse

import "strconv"

// Int returns its integer argument with its decimal digits reversed,
// preserving the sign.
func Int(i int) int {
	neg := i < 0
	if neg {
		i = -i
	}
	n, _ := strconv.Atoi(String(strconv.Itoa(i)))
	if neg {
		n = -n
	}
	return n
}
