package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var field [9][9]int

func initField() {
	readField()
	printField()
}

func readField() {
	jsonFile, err := os.Open("field.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened field.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &field)
}

func printField() {
	fmt.Println("---==[ CURRENT FIELD ]==---")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if field[i][j] == 0 {
				fmt.Printf("| ")
			} else {
				fmt.Printf("|%d", field[i][j])
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\n\n")
}

func main() {
	start := time.Now()

	initField()

	solve()
	printField()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
