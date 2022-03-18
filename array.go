package main

import "fmt"

func main() {
	fmt.Println("welcome array")

	var names [3]string

	names[0] = "jelly"
	names[1] = "jelly1"
	names[2] = "jelly2"

	fmt.Println(names)

	var values = []int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(values)

}
