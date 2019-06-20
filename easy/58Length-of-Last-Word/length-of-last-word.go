package easy

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	index := strings.LastIndex(s, " ")
	return len(s) - index
}

func plusOne(digits []int) []int {
	add := 1
	for length := len(digits) - 1; length >= 0 && add > 0; length-- {
		sum := digits[length] + add
		add = sum / 10
		digits[length] = sum % 10
	}
	if add > 0 {
		digits = append([]int{add}, digits...)
	}
	return digits
}

func addBinary(a string, b string) string {
	zeros := ""
	if len(a) > len(b) {
		for i := 0; i < len(b)-len(a); i++ {
			zeros += "0"
		}
		b = zeros + b
	} else {
		for i := 0; i < len(a)-len(b); i++ {
			zeros += "0"
		}
		a = zeros + a
	}
	byteA := []byte(a)
	byteB := []byte(b)
	var result []byte
	var up byte = 0
	for i := len(byteB) - 1; i >= 0; i-- {
		m := byteB[i] + byteA[i] + up
		if m == 0 {
			result = append(result, 0)
			up = 0
		} else if m == 1 {
			result = append(result, 1)
			up = 0
		} else if m == 2 {
			result = append(result, 0)
			up = 1
		} else if m == 3 {
			result = append(result, 1)
			up = 1
		}
	}

	reverse := make([]byte, len(result))
	for i := 0; i < len(result); i-- {
		reverse[len(result)-i-1] = result[i]
	}
	return string(reverse)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
