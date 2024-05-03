package twosum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]

		v, ok := numMap[c]
		if ok {
			return []int{v, i}
		}

		numMap[nums[i]] = i
	}

	return nil
}
