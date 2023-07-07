package main

// check if field is solved or if there are any illegal placements
// return: errors, solved
func checkField() (bool, bool) {
	solved := true
	// check all rows
	for i := 0; i < 9; i++ {
		var m = make(map[int]bool)
		for j := 0; j < 9; j++ {
			d := field[i][j]
			if d == 0 {
				solved = false
			}
			if m[d] == true {
				return false, false
			}
			m[d] = true
		}
	}

	// check all cols
	for i := 0; i < 9; i++ {
		var m map[int]bool
		for j := 0; j < 9; j++ {
			d := field[j][i]
			if d == 0 {
				solved = false
			}
			if m[d] == true {
				return false, false
			}
			m[d] = true
		}
	}

	// check squares
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			// check square (x,y)
			var m map[int]bool
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					d := field[x*3+i][y*3+j]
					if d == 0 {
						solved = false
					}
					if m[d] == true {
						return false, false
					}
					m[d] = true
				}
			}

		}
	}

	return true, solved
}
