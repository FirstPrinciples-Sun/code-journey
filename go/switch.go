package main

import "fmt"

func main() {
	a := 1

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	switch {
	case a < 1:
		fmt.Println("less than one")
		fallthrough
	case a > 2:
		fmt.Println("Mor than one")
		fallthrough
	default:
		fmt.Println("default")                                      
	}
}