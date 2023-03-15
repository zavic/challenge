package main

import (
	"fmt"
	"os"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	args := os.Args
	p := Person{
		Nama:      args[1],
		Alamat:    args[2],
		Pekerjaan: args[3],
		Alasan:    args[4],
	}
	perkenalan(p)
}

func perkenalan(p Person) {
	fmt.Println("Nama: ", p.Nama)
	fmt.Println("Alamat: ", p.Alamat)
	fmt.Println("Pekerjaan: ", p.Pekerjaan)
	fmt.Println("Alasan: ", p.Alasan)
}
