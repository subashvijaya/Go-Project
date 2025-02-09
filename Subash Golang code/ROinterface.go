package main

import "fmt"

type redbustravel interface {
	redbus()
}

type orangebustravel interface {
	orangebus()
}

type person struct {
	fname       string
	lname       string
	age         int
	place       string
	pickuppoint string
	droppoint   string
}

func (r person) redbus() {
	fmt.Println(r.fname, r.lname, r.age, r.place, r.pickuppoint, r.droppoint)

}

func (r person) orangebus() {
	fmt.Println(r.fname, r.lname, r.age, r.place, r.pickuppoint, r.droppoint)

}

func main() {

	var trip person = person{"v", "suba", 24, "banglore", "HAL", "KRpuram"}
	var det redbustravel = trip
	det.redbus()
	var trip1 person = person{"s", "ss", 22, "sankai", "lnp", "trp"}
	var det1 orangebustravel = trip1
	det1.orangebus()

}
