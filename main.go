package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func initField() *Field {
	f := readField()
	f.printField()
	return f
}

type Field struct {
	field [9][9]int
}

func NewField(f [9][9]int) *Field {
	return &Field{field: f}
}

func readField() *Field {
	jsonFile, err := os.Open("field.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened field.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var field [9][9]int

	json.Unmarshal(byteValue, &field)

	return NewField(field)
}

func (f *Field) printField() {
	fmt.Println("---==[ CURRENT FIELD ]==---")
	for i := 0; i < 9; i++ {
		for j := 0; j <= 9; j++ {
			if i%3 == 0 && j == 0 {
				fmt.Print("┏")
			} else if (i+1)%3 == 0 && j == 0 {
				fmt.Print("┗")
			} else if i%3 == 0 && j == 9 {
				fmt.Print("┓")
			} else if (i+1)%3 == 0 && j == 9 {
				fmt.Print("┛")
			} else if i%3 == 0 && (j == 3 || j == 6) {
				fmt.Print("┃")
			} else if (i+1)%3 == 0 && (j == 3 || j == 6) {
				fmt.Print("┃")
			} else if j%3 == 0 {
				fmt.Print("┃")
			} else {
				fmt.Print("│")
			}

			if j < 9 {
				if f.field[i][j] == 0 {
					fmt.Printf(" ")
				} else {
					fmt.Printf("%d", f.field[i][j])
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func main() {
	start := time.Now()

	f := initField()

	d := false
	f.field, d = solve(*f)
	if d {
		fmt.Println("solving failed")
	}
	f.printField()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
