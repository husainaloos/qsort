package main

import (
	"reflect"
	"testing"
)

func TestHoareScheme(t *testing.T) {
	tc := [][]int{
		[]int{5, 4, 3, 2, 1},
		[]int{3, 5, 4, 1, 2},
		[]int{1, 2, 3, 4, 5},
	}
	expected := []int{1, 2, 3, 4, 5}
	for _, arr := range tc {
		scheme := HoareScheme{}
		scheme.Sort(arr)
		if !reflect.DeepEqual(expected, arr) {
			t.Errorf("expected %v, found %v", expected, arr)
		}
	}

}
