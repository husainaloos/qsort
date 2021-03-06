package main

type HoareScheme struct{}

func (scheme *HoareScheme) Sort(arr []int) {
	scheme.qsort(arr, 0, len(arr)-1)
}

func (scheme *HoareScheme) qsort(arr []int, low, high int) {
	if low < high {
		p := scheme.partition(arr, low, high)
		scheme.qsort(arr, low, p)
		scheme.qsort(arr, p+1, high)
	}
}

func (scheme *HoareScheme) partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1
	for {
		i++
		for arr[i] < pivot {
			i++
		}
		j--
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		scheme.swap(arr, i, j)
	}
	return -1
}

func (scheme *HoareScheme) swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
