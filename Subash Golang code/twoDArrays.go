package main

import "fmt"

func main() {

	//var rows int
	//var cols int

	var twoDArray [3][4]int

	fmt.Println("Two-dimensional array...!")

	twoDArray = [3][4]int{
		{2, 4, 6, 8},
		{1, 3, 7, 8},
		{1, 2, 3, 5},
	}

	for i := 0; i < 3; i++ {

		for j := 0; j < 4; j++ {

			fmt.Printf("%d\t", twoDArray[i][j])
		}
		fmt.Println()
	}
}
