package main

import "fmt"

// check if field is solved or if there are any illegal placements
// return: errors, missing (0 == solved)
func (f *Field) checkField() (bool, int) {
	missing := 0
	err := false

	// check all rows
	for i := 0; i < 9; i++ {
		var m = make(map[int]bool)
		for j := 0; j < 9; j++ {
			d := f.field[i][j]
			if d == 0 {
				missing++
				continue
			}
			if m[d] == true {
				err = true
			}
			m[d] = true
		}
	}

	// check all cols
	for i := 0; i < 9; i++ {
		var m = make(map[int]bool)
		for j := 0; j < 9; j++ {
			d := f.field[j][i]
			if d == 0 {
				continue
			}
			if m[d] == true {
				err = true
			}
			m[d] = true
		}
	}

	// check squares
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			// check square (x,y)
			var m = make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					d := f.field[x*3+i][y*3+j]
					if d == 0 {
						continue
					}
					if m[d] == true {
						err = true
					}
					m[d] = true
				}
			}

		}
	}
	fmt.Printf("missing: %d\n", missing)
	if err {
		fmt.Printf("has errors\n")
	}
	if missing == 0 {
		fmt.Printf("is solved\n")
	}
	return err, missing
}

func (f *Field) possibleNumbers(x int, y int) []int {
	var m = make(map[int]bool)
	for i := 1; i <= 9; i++ {
		m[i] = true
	}

	for i := 0; i < 9; i++ {
		// check row
		if i != y {
			d := f.field[x][i]
			m[d] = false
		}
		// check col
		if i != x {
			d := f.field[i][y]
			m[d] = false
		}
	}

	x1 := x / 3
	y1 := y / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			d := f.field[x1*3+i][y1*3+j]
			m[d] = false
		}
	}

	var r []int
	for i := 1; i <= 9; i++ {
		if m[i] {
			r = append(r, i)
		}
	}
	return r
}
