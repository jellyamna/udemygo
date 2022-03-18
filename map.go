package main

import (
	"encoding/json"
	"fmt"
)

type Sample struct {
	name string
}

func main() {
	fmt.Println("string"[0])
	var mapkosong = map[string]string{}

	// var mapkosongtest = map[string]string{}
	mapkosongtest3 := map[string]string{}
	var sample []Sample
	sample1 := make([]Sample, 0, 0) //cara satu
	sample2 := []Sample{}           //cara dua

	fmt.Println("sample1", len(sample1))
	fmt.Println("sample2", len(sample2))

	// mapkosongtest2 = make(map[string]string)

	newmap := map[string]string{

		"name":   "jelly",
		"alamat": "pamulang",
	}

	fmt.Println(newmap)
	fmt.Println(mapkosong)

	mapkosong["title"] = "programmer"
	fmt.Println(mapkosong)
	fmt.Println(mapkosong["title"])

	fmt.Println("len", len(mapkosongtest3))
	jsonStr, _ := json.Marshal(sample)
	jsonStr1, _ := json.Marshal(sample1)

	/**
	dokumentasi kan funtion
	*/
	fmt.Println(string(jsonStr))
	fmt.Println(string(jsonStr1))
}
