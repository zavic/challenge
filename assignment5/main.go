package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func kondisi1() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	interface1 := []string{"bisa1", "bisa2", "bisa3"}
	interface2 := []string{"coba1", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(interface1, i+1)
		}(i)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(interface2, i+1)
		}(i)
	}
	wg.Wait()
}

func kondisi2() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	interface1 := []string{"bisa1", "bisa2", "bisa3"}
	interface2 := []string{"coba1", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println(interface1, i+1)
		}(i)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println(interface2, i+1)
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Kondisi 1")
	kondisi1()
	fmt.Println("\n======================\nKondisi 2")
	kondisi2()
}
