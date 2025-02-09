package main

import "fmt"

func main() {

	fmt.Println("find the given array odd and even number...!")

	arr := []int{12, 3, 4, 7, 11, 16, 12, 1, 4, 2}

	for i := 0; i < len(arr); i++ {

		if arr[i]%2 == 0 {

			fmt.Println("even number", arr[i])
		} else {
			fmt.Println("odd number", arr[i])
		}
	}

}
