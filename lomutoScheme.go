package main

type LomutoScheme struct {
}

func (scheme *LomutoScheme) Sort(arr []int) {
	scheme.qsort(arr, 0, len(arr)-1)
}

func (scheme *LomutoScheme) qsort(arr []int, low, high int) {
	if low < high {
		p := scheme.partition(arr, low, high)
		scheme.qsort(arr, low, p-1)
		scheme.qsort(arr, p+1, high)
	}
}

func (scheme *LomutoScheme) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			if i != j {
				i++
				scheme.swap(arr, i, j)
			}
		}
	}
	i++
	scheme.swap(arr, i, high)
	return i
}

func (scheme *LomutoScheme) swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
