package main

import "fmt"

type jenkins string

func (j jenkins) Greet() {
	fmt.Println("jenkins plugin")
}

var Jenkins jenkins
