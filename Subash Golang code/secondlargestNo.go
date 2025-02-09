package main

import "fmt"

func main() {

	arr := [5]int{3, 5, 15, 3, 9}

	var firstlarge, secondlarge int

	for _, num := range arr {

		if num > firstlarge {

			secondlarge = firstlarge
			firstlarge = num
		} else if num > secondlarge && num != firstlarge {
			secondlarge = num
		}
	}
	fmt.Println("The given array second largest number is :", secondlarge)
}
