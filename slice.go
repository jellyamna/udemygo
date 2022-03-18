package main

import "fmt"

func main() {

	var monnths = [...]string{
		"januari",   //0
		"februari",  //1
		"maret",     //2
		"april",     //3
		"may",       //4
		"juni",      //5
		"juli",      //6
		"agustus",   //7
		"september", //8
		"oktober",   //9
		"november",  //10
		"desember",  //11
	}

	//membuat slice dari array
	slice1 := monnths[4:7]
	slice2 := monnths[4:]
	slicenew := monnths[10:]

	slice3 := monnths[:11]

	fmt.Println(slicenew)
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8
	fmt.Println("monnths[4:]", slice2)

	fmt.Println("monnths[:11]", slice3)

	//append

	slicenew[0] = "november_new"
	fmt.Println(monnths)
	slicenew = append(slicenew, "bulan_baru")
	fmt.Println("cetak append")
	fmt.Println(monnths)
	fmt.Println(slicenew)

}
