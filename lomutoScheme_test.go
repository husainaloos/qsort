package main

import (
	"bufio"
	"os"
	"path"
	"reflect"
	"strconv"
	"testing"
)

func TestLomutoScheme(t *testing.T) {
	tc := [][]int{
		[]int{5, 4, 3, 2, 1},
		[]int{3, 5, 4, 1, 2},
		[]int{1, 2, 3, 4, 5},
	}
	expected := []int{1, 2, 3, 4, 5}
	for _, arr := range tc {
		scheme := LomutoScheme{}
		scheme.Sort(arr)
		if !reflect.DeepEqual(expected, arr) {
			t.Errorf("expected %v, found %v", expected, arr)
		}
	}

}

func BenchmarkLomutoScheme_1000(b *testing.B) {
	for fn := 0; fn < 100; fn++ {
		filename := "arr_" + strconv.Itoa(fn) + ".txt"
		filepath := path.Join("./data/arr_1000", filename)
		arr, err := readFile(filepath)
		if err != nil {
			b.Error("cannot read file ", filepath, err)
		}
		b.Run("file "+filepath, func(b *testing.B) {
			scheme := LomutoScheme{}
			for i := 0; i < b.N; i++ {
				scheme.Sort(arr)
			}
		})

	}
}
func BenchmarkLomutoScheme_10000(b *testing.B) {
	for fn := 0; fn < 100; fn++ {
		filename := "arr_" + strconv.Itoa(fn) + ".txt"
		filepath := path.Join("./data/arr_10000", filename)
		arr, err := readFile(filepath)
		if err != nil {
			b.Error("cannot read file ", filepath, err)
		}
		b.Run("file "+filepath, func(b *testing.B) {
			scheme := LomutoScheme{}
			for i := 0; i < b.N; i++ {
				scheme.Sort(arr)
			}
		})

	}
}

/*func BenchmarkLomutoScheme_100000(b *testing.B) {
	for fn := 0; fn < 100; fn++ {
		filename := "arr_" + strconv.Itoa(fn) + ".txt"
		filepath := path.Join("./data/arr_100000", filename)
		arr, err := readFile(filepath)
		if err != nil {
			b.Error("cannot read file ", filepath, err)
		}
		b.Run("file "+filepath, func(b *testing.B) {
			scheme := LomutoScheme{}
			for i := 0; i < b.N; i++ {
				scheme.Sort(arr)
			}
		})

	}
}
func BenchmarkLomutoScheme_1000000(b *testing.B) {
	for fn := 0; fn < 100; fn++ {
		filename := "arr_" + strconv.Itoa(fn) + ".txt"
		filepath := path.Join("./data/arr_1000000", filename)
		arr, err := readFile(filepath)
		if err != nil {
			b.Error("cannot read file ", filepath, err)
		}
		b.Run("file "+filepath, func(b *testing.B) {
			scheme := LomutoScheme{}
			for i := 0; i < b.N; i++ {
				scheme.Sort(arr)
			}
		})

	}
}
*/
func readFile(filepath string) ([]int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	arr := make([]int, 1000)
	for scanner.Scan() {
		str := scanner.Text()
		i, _ := strconv.Atoi(str)
		arr = append(arr, i)
	}
	return arr, nil
}
