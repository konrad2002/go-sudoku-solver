package main

import (
	"fmt"
	"testing"
)

func Test_possibleNumbers(t *testing.T) {
	initField()
	a := possibleNumbers(0, 2)
	fmt.Println(a)
}
