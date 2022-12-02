package aaa

func twoSum(nums []int, target int) []int {
	maps := map[int]int{}
	num := len(nums)
	for i := 0; i < num; i++ {
		if x, ok := maps[target-nums[i]]; ok {
			return []int{i, x}
		}
		maps[nums[i]] = i
	}
	return nil
}
