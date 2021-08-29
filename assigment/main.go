package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range n {
		if v%2 == 0 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
