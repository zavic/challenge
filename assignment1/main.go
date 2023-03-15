package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true
	var ya rune = 'Я'
	var k float64 = 123.456

	fmt.Printf("%v\n", i)   // Menampilkan nilai var i
	fmt.Printf("%T\n", i)   // Menampilkan tipe data var i
	fmt.Printf("%v\n", "%") // Menampilkan tanda %
	fmt.Printf("%t\n\n", j) // Menampilkan nilai boolean j

	fmt.Printf("%b\n", i)    // Menampilkan nilai base 2 var i
	fmt.Printf("%c\n", ya)   // Menampilkan unicode russia Я
	fmt.Printf("%d\n", i)    // Menampilkan nilai base 10 var i
	fmt.Printf("%o\n", i)    // Menampilkan nilai base 8 var i
	fmt.Printf("%x\n", 15)   // menampilkan base 16 : f
	fmt.Printf("%X\n", 15)   // menampilkan base 16 : F
	fmt.Printf("%U\n\n", ya) // menampilkan unicode karakter Я : U+042F

	fmt.Printf("%f\n", k) // menampilkan float : (123.456000)
	fmt.Printf("%E\n", k) // menampilkan float scientific : (1.234560E+02)
}
