package homeworks

func SmallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		count := 0
		for j := range nums {
			if nums[i] > nums[j] {
				
			}
		}
		result[i] = count
	}
	return result
}
