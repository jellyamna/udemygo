package main

import "fmt"

type Name struct {
	NamaDepan    string
	NamaBelakang string
}

func main() {

	var name Name

	name.NamaDepan = "jelly"
	name.NamaBelakang = "amna"

	name1 := name
	name2 := &name
	name1.NamaDepan = "kevin"

	var name3 *Name
	var name4 Name

	*name2 = Name{
		NamaDepan: "tasya",
	}

	fmt.Println(name1)
	fmt.Println(name)
	fmt.Println(*name2)
	fmt.Println(name3)
	fmt.Println(name4)

}
