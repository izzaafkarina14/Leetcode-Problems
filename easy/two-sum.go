package easy

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		index1 := target - nums[i]
		for n := i + 1; n < len(nums); n++ {
			if index1 == nums[n] {
				return []int{i, n}
			}
		}
	}
	return []int{}
}