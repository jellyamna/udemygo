package main

import "fmt"

func main() {

	var nilai1 uint32 = 343
	var name string = "jelly"
	fmt.Println("nilai variable nilai1 uint64 = ", uint64(nilai1))
	fmt.Println("nilai variable nilai1 int64 = ", int64(nilai1))
	var e byte //alias dari uint8

	e = name[4]
	fmt.Println("karakter ke - 1 =:", string(name[1]))
	fmt.Println(e)
	fmt.Println(string(e))
	fmt.Println("y"[0])

}
