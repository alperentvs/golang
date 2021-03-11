// Package clause
package main

// Import statement
import "fmt"

// My Code
func main() {

	// var - name of variable - static type

	/* var name string
	   name = "Alperen"
	*/

	// name := "Alperen"  -- bu da farklı bir gösterim.

	/* var age int
	   age = 24
	*/

	// age := 24

	/* var isMarried bool
	   isMarried = true
	*/

	// isMarried := true

	var name string = "Alperen"

	var firstName, lastName string = "Alperen", "Tavas"

	var age int = 24

	var isMarried bool = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(firstName, lastName)
}
