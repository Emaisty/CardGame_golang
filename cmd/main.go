package main

import "fmt"

type lol struct {
	name string
}

func kke() *lol {
	a := lol{name: "kek"}
	var b *lol
	b = &a
	return b
}

func main() {
	a := kke()
	fmt.Println(a.name)
}
