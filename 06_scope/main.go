package main

import "fmt"

/* var packVar = "Package Scope"
var funcVar = "Func(Package) Scope" */

func main() {

	/* if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	} */

	/* var funcVar = "Func Scope"
	fmt.Println(funcVar) */

	var name = "Mehmet"
	name, surname := "Alperen", "Tavas" // "Surname" değişkeni oluşturmasaydık kısa gösterim yapamazdık.
	fmt.Println(name, surname)

	/* 	fmt.Println(packVar)

	   	myFunc() */

}

/* func myFunc() {
	fmt.Println(funcVar)
} */
