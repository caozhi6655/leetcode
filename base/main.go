package main

import "fmt"

func main() {
	arr := []int{0, 5, 2, 6, 4, 1, 3, 17, 9, 13}
	sort := mergeSort(arr)
	fmt.Println(sort)
}

//冒泡排序
func bubbleSort(arr []int) {
	for out := len(arr) - 1; out > 0; out-- {
		for in := 0; in < out; in++ {
			if arr[in] > arr[out] {
				arr[in], arr[out] = arr[out], arr[in]
			}
		}
	}
}

//选择排序
func selectSort(arr []int) {
	min := 0

	for out := 0; out < len(arr)-1; out++ {
		min = out
		for in := out + 1; in < len(arr); in++ {
			if arr[in] < arr[min] {
				min = in
			}
		}
		arr[min], arr[out] = arr[out], arr[min]
	}
}

//插入排序
func insertSort(arr []int) {
	for out := 1; out < len(arr); out++ {
		in := out
		tmp := arr[out]
		for in > 0 && arr[in-1] > tmp {
			arr[in] = arr[in-1]
			in--
		}
		arr[in] = tmp
	}
}

func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	partitionIndex := partition(arr, left, right)
	_quickSort(arr, left, partitionIndex-1)
	_quickSort(arr, partitionIndex+1, right)
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	for {
		for {
			if right > 0 && arr[right] > pivot {
				right--
			} else {
				break
			}

		}
		for {
			if arr[left] < pivot {
				left++
			} else {
				break
			}
		}
		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left], arr[right] = arr[right], arr[left]
	return right
}

//归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	return merge(mergeSort(arr[0:middle]), mergeSort(arr[middle:]))
}

func merge(left []int, right []int) []int {
	var result []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}

	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}
