package main

import "fmt"

func main() {

	myperson := map[string]int{"sg": 19, "ss": 21, "sn": 23} //creating a map

	fmt.Println(myperson)

	keytoupdate := "ss" //update map
	if value, ex := myperson[keytoupdate]; ex {
		myperson[keytoupdate] = value + 1
		fmt.Printf("update value for key : %s : %d\n", keytoupdate, myperson[keytoupdate])
	} else {
		fmt.Printf("Key '%s' not found in the map\n", keytoupdate)
	}

	//keytodelete map
	for v := range myperson {

		if v == "sn" {
			delete(myperson, v)
		}
	}
	fmt.Println(myperson)

	keytochange := "sg" //updateto chane map
	if value, ex := myperson[keytochange]; ex {
		myperson[keytochange] = value * 2
		fmt.Printf("update value for key : %s : %d\n", keytochange, myperson[keytochange])
	} else {
		fmt.Printf("Key '%s' not found in the map\n", keytochange)
	}

	fmt.Println("map after change : ", myperson)

}

/*package main

import "fmt"

func delete_key(p map[string]int){

    for v := range person{

        if v == "ram"{
            delete(p,v)
        }
    }
}

func main(){

    person := make(mape[string]int)

    person["raja"] = 99
    person["ram"] = 95
    person["jeevan"] = 100
    person["suba"] = 85

    fmt.Printn(person)
    delete_key(person)
}*/
