package main

import "fmt"

func main() {
	arr := []int{0, 5, 2, 6, 4, 1, 3, 17, 9, 13}
	sort := mergeSort(arr)
	fmt.Println(sort)
}

func bubbleSort(arr []int) {
	for out := len(arr) - 1; out > 0; out-- {
		for in := 0; in < out; in++ {
			if arr[in] > arr[out] {
				arr[in], arr[out] = arr[out], arr[in]
			}
		}
	}
}

func selectSort(arr []int) {
	for out := 0; out < len(arr)-1; out++ {
		min := out
		for in := out + 1; in < len(arr); in++ {
			if arr[in] < arr[min] {
				min = in
			}
		}
		arr[out], arr[min] = arr[min], arr[out]
	}
}

func insertSort(arr []int) {
	for out := 1; out < len(arr); out++ {
		tmp := arr[out]
		in := out
		for in > 0 && arr[in-1] > tmp {
			arr[in] = arr[in-1]
			in--
		}
		arr[in] = tmp
	}
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	partition := partitionIt(arr)
	quickSort(arr[:partition])
	quickSort(arr[partition:])
}

func partitionIt(arr []int) int {
	left := 0
	right := len(arr) - 1
	pivot := arr[right]

	for {
		for {
			if arr[left] < pivot {
				left++
			} else {
				break
			}
		}

		for right > 0 {
			if arr[right] > pivot {
				right--
			} else {
				break
			}
		}

		if left >= right {
			break
		}

		arr[left], arr[right] = arr[right], arr[left]
	}

	return left
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}

	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
