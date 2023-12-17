package main

import "fmt"

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
	fmt.Println(f.Bar(99))
	fmt.Println(f.Bar(101))
	fmt.Println(f.Bar(1))
}
