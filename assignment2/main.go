package main

import "fmt"

func main() {
	i := 0
	for i < 5 {
		fmt.Println("Nilai i =", i)
		i++
		if i == 5 {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					input := "САЦАРВО"
					for index, value := range input {
						fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index)
					}
				} else {
					fmt.Printf("Nilai j = %d\n", j)
				}
			}
		} else {
			continue
		}
	}
}
