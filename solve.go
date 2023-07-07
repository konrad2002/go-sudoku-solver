package main

func solve() {
	e, s := checkField()
	if e {
		panic("error")
	}
	if s {
		panic("solved")
	}
}
