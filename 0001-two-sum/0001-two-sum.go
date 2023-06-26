func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, f := m[target - num]; f {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil
}