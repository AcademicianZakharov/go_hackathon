package main

import (
	"fmt"
	//"errors"
)

func main() {
	var intarr [3]int32 = [...]int32{1, 2, 3}
	fmt.Println(intarr[1:3])
	fmt.Println(len(intarr))

	var intslice []int32 = []int32{4, 5, 6}
	fmt.Println(intslice)
	intslice = append(intslice, 7, intarr[0])
	fmt.Println(intslice)

	//using a function from another file
	printSomething("\n---------gabriels test section--------------")
	// var mymap0 map[string]int = make(map[string]int)
	// fmt.Println(mymap0)

	// var mymap = map[string]int{"sam": 23, "sarah": 45}
	// fmt.Println(mymap)
	// var age, ok = mymap["sam"]
	// fmt.Println(age, ok)
	// //delete(mymap, "sarah")

	// for name, age := range mymap {
	// 	//fmt.Printf("Name %v, age %v\n", name, mymap[name])
	// 	fmt.Printf("Name %v, age %v\n", name, age)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	var mystr = []rune("text")
	var indexed = mystr[0]
	fmt.Println("%v, %T", indexed, mystr)

}
