package main

import "fmt"

func main() {

	rows := 5

	//Logic : 1

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			fmt.Printf("& ")
		}
		fmt.Println()
	}
	fmt.Println("\n")

	//Logic : 2
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("$ ")
		}
		fmt.Println()
	}
	fmt.Println("\n")

	//Logic : 3
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows+1-i; j++ {
			fmt.Printf("$ ")
		}
		fmt.Println()
	}
	fmt.Println("\n")

	//Logic : 4
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println("\n")

	//Logic : 5
	for i := 1; i <= rows; i++ {
		for space := 0; space <= rows-i; space++ {
			fmt.Println(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		for j := 2; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println("\n")

	//logic : 6
	var i, j, k, alphabet int

	for i = 0; i < rows; i++ {
		alphabet = 65
		for j = rows; j > i; j-- {
			fmt.Printf(" ")
		}
		for k = 0; k <= i; k++ {
			fmt.Printf("%c ", alphabet+k)
		}

		fmt.Println()
	}
}


