package main

import "fmt"

func solve(f Field) ([9][9]int, bool) {
	fmt.Printf("\nSOLVE FUNCTION has been called\n")
	_, m1 := f.checkField()
	if m1 == 0 {
		return f.field, true
	}
	p := 10 // minimum number of possible options in all fields
	fmt.Println("searching field with 1 possible number...")
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if f.field[x][y] != 0 {
				continue
			}
			a := f.possibleNumbers(x, y)
			p1 := len(a)
			if p1 < p {
				p = p1
			}
			if p1 == 1 {
				fmt.Printf("1 option field found: (%d,%d)\n", x, y)
				fmt.Printf("[+] adding number '%d'\n", a[0])
				f.field[x][y] = a[0]
				// no check for errors needed since possibleNumbers sticks to the rules
				return solve(f)
			}
		}
	}

	// if no field with only one option occurs, continue with more
	fmt.Printf("no 1 option field found, contiue looking for %d options (min)...\n", p)
	// try both options for fields with p options
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if f.field[x][y] != 0 {
				continue
			}
			a := f.possibleNumbers(x, y)
			if len(a) == p {
				fmt.Printf("found field with %d options at (%d,%d), testing all options:\n", p, x, y)
				for i := 0; i < p; i++ {
					fmt.Printf(" -> (%d,%d) option %d: %d\n", x, y, i, a[i])
					f.field[x][y] = a[i]
					f1, done := solve(f)
					if done {
						return f1, true
					}
				}
				// no of the options was right -> unsolvable
				fmt.Println("solving the sudoku with the current setting of numbers is impossible, returning to point of last decision")
				return f.field, false
			}
		}
	}
	// no of the options was right -> unsolvable
	return f.field, false
}
