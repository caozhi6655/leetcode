package easy

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0

	for j := 0; j < len(nums); j++ {
		if val != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
