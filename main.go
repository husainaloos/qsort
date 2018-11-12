package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxSize := 10
	for size := 1; size <= maxSize; size++ {
		arr := generateRandomArray(size)
		scheme := LomutoScheme{}
		scheme.Sort(arr)
		fmt.Println(arr)
	}

}

func generateRandomArray(size int) []int {
	rand := rand.New(rand.NewSource((int64)(time.Now().Nanosecond())))
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}
