package main

import (
// "log"
)

func find(big, small [][]int) (int, int) {
	match := false
	sm_width := len(small)
	sm_length := len(small[0])

	x_ref := -1
	y_ref := -1

	for i := range big {
		for j := range big[i] {
			if !match {
				// if not in middle of another match
				if big[i][j] == small[0][0] {
					match = true
					x_ref = i
					y_ref = j
				}
			} else if i-x_ref >= 0 && i-x_ref < sm_width && j-y_ref >= 0 && j-y_ref < sm_length {
				// if in the middle of another match, need to make sure
				// that we havent extended beyond width of smaller image
				if big[i][j] != small[i-x_ref][j-y_ref] {
					// if in the middle of match, still within limits of smaller photo
					// and found a mismatching pixel then reset
					match = false
					x_ref = -1
					y_ref = -1
				}
			}
			if match && i-x_ref >= sm_width && i-y_ref >= sm_length {
				// once we know we considered all pixels of smaller photo and all pass
				// then stop looking and report coordiantes

				return x_ref, y_ref
			}

		}
	}
	return x_ref, y_ref
}

// iterate over big image, row first, col next
// if
