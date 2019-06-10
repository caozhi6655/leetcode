package easy

// SearchInsert 插入
func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums) - 1

	for {
		mid := (high-low)/2 + low

		if nums[mid] == target {
			return mid
		} else if high-low == 0 {
			if nums[low] < target {
				return low + 1
			}

			return low

		} else if high-low > 0 && nums[mid] > target {
			high = mid
		} else if high-low > 0 && nums[mid] < target {
			low = mid + 1
		}
	}
}

func SearchInsert1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}
