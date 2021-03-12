// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluştuurp, çıktısını yazdırınız.

/* package main

import "fmt"

func main() { */

/* var studentName string = "John Doe"
var grade int = 77
var isPassed bool = true */

/* var studentName = "John Doe"
var grade = 77
var isPassed = true */

/* 	studentName := "John Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

} */

// 2-) Yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.

/* package main

import "fmt"

func main() {

	studentName, grade, isPassed := "Jonh Doe", 77, true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

} */

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

/* package main

import "fmt"

func main() {

	// var studentName string   // --> Bu aşama "Declaration" aşamasıdır.
	studentName := "Alperen Tavas" // "Assign" aşaması atama yaptığımzı kısımdır.
	// var studentName string = "John Doe" --> Bu aşama "Initialization" aşamasıdır. Bir değeri oluşturup başlatma gibi düşünebiliriz.
	// "Initial Value" ilk değer gibi düşünebiliriz. studentName değişkenine "John Doe" yi atadığımız gibi.

	fmt.Println(studentName)

} */

// 4-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

package main

import "fmt"

func main() {

	/* var studentName string = "John Doe"
	studentName = "Alperen Tavas" */

	studentName := "John Doe"
	studentName = "Alperen Tavas"

	fmt.Println(studentName)

}
