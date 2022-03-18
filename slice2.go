package main

import "fmt"

type Data struct {
	nama string
}

func main() {

	slice1 := make([]string, 2, 5)

	var data []Data

	var sdata Data
	sdata.nama = "dfdfd"
	data = append(data, sdata)

	slice1[0] = "jelly"
	slice1[1] = "amna"
	fmt.Println(slice1)

	if data == nil {
		fmt.Println("data nil")
	}
	fmt.Println(data)

}
