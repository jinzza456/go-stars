package main

import "fmt"

func main() {

	for i := 1; i < 5; i++ {
		for j := 1; j < 5-i; j++ {
			fmt.Print(" ")
		}
		for x := 1; x <= (i*2 - 1); x++ {
			fmt.Print("*")
		}
		println("")
	}
	for i := 1; i < 4; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Print(" ")
		}
		for x := 1; x <= (7 - i*2); x++ {
			fmt.Print("*")
		}
		println("")
	}
}
