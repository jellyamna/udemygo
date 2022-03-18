package main

import "fmt"

func main() {

	i := 80
	var floatdata float64 = float64(i)
	var data int = abc(10).(int)
	fmt.Println(abc(10))
	fmt.Println(data)
	fmt.Println(floatdata)

}

func abc(i int) interface{} {
	return i
}
