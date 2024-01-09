package main

import "fmt"

func quickSort(arr []int, l, r int) {
	if l < r {
		pivot := partition(arr, l, r)
		quickSort(arr, l, pivot-1)
		quickSort(arr, pivot+1, r)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	i := l + 1
	j := r

	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for j >= i && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j,j,j,j,,j,j,j,j
}
func main() {
	arr := []int{9, 5, 2, 7, 1, 8, 3}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
