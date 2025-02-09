package main

import "fmt"

func Threesum(num [7]int, target int) []int {

	res := []int{}

	for i := 0; i < len(num); i++ {

		for j := 1; j < len(num); j++ {

			for k := 2; k < len(num); k++ {

				if num[i]+num[j]+num[k] == target && i != j && j != k {
					res = []int{i, j, k}
					return res
				}
			}
		}
	}
	return res
}

func main() {

	fmt.Println("Threesum Leetcode Problem")

	num := [7]int{2, 7, 8, 11, 15, 13, 85}

	target := 100

	res := Threesum(num, target)

	fmt.Println(res)
}
