package trappingrainwater

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]

	res := 0

	for left != right {
		if height[left] < height[right] {
			left++

			if height[left] > maxLeft {
				maxLeft = height[left]
			}

			res += maxLeft - height[left]
		} else {
			right--

			if height[right] > maxRight {
				maxRight = height[right]
			}

			res += maxRight - height[right]
		}
	}

	return res
}
