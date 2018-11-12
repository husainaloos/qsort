package main

import (
	"fmt"
	_ "gonum.org/v1/plot"
	_ "gonum.org/v1/plot/plotter"
	"math/rand"
	"time"
)

type coordinate struct {
	x, y float64
}

type coordinates []coordinate

func (c coordinates) Len() int {
	return len(c)
}

func (c coordinates) XY(i int) (x, y float64) {
	return c[i].x, c[i].y
}

func main() {
	maxSize := 10

	for size := 1; size <= maxSize; size++ {
		arr := generateRandomArray(size)
		scheme := LomutoScheme{}
		t := time.Now()
		scheme.Sort(arr)
		delta := time.Since(t)
		fmt.Println(delta)
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
