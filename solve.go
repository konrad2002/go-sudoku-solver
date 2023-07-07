package main

import "fmt"

func solve() {
	checkField()
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			a := possibleNumbers(x, y)
			fmt.Println(a)
		}
	}
}
