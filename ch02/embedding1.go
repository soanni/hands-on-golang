package main

import "fmt"

type Bird struct {
	featherLength int
	classification string
}

type Pigeon struct {
	Bird
	featherLength float64
	Name string
}

func main() {
	p := Pigeon{Name: "Tweety",}
	p.Bird.featherLength = 314
	fmt.Println(p)
}
