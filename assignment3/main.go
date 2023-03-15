package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "selamat malam"
	strSlice := strings.Split(str, "")

	for _, char := range strSlice {
		fmt.Println(char)
	}

	charCountMap := make(map[string]int)
	for _, char := range strSlice {
		charCountMap[char]++
	}

	fmt.Println(charCountMap)
}
