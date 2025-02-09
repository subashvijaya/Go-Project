package main

import "fmt"

func main() {

	var eventotal int

	var oddtotal int

	fmt.Print("Number 1 to 100 find Even and Odd Sum\n ")

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {

			eventotal += i
			fmt.Printf("%d\t", i)

		} else {
			oddtotal += i
		}
	}
	fmt.Println("\nSum of Even Numbers from 1 to 100 = ", eventotal)
	fmt.Println("\nSum of Odd Numbers from 1 to 100  =  ", oddtotal)

}
