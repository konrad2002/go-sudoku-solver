package main

func solve() {
	checkField()
	for i := 0; i < 810; i++ {
		for x := 0; x < 9; x++ {
			for y := 0; y < 9; y++ {
				if field[x][y] != 0 {
					continue
				}
				a := possibleNumbers(x, y)
				if len(a) == 1 {
					field[x][y] = a[0]
					_, miss := checkField()
					if miss == 0 {
						return
					}
				}
			}
		}
	}

}
