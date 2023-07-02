func runningSum(nums []int) []int {
    acc := nums[0]
    temp := []int{acc}
    for i:= 1; i < len(nums); i++ {
        acc += nums[i]
        temp = append(temp, acc)
    }

    return temp
}