func maximumWealth(accounts [][]int) int {
    var maxNum float64
    for i := range accounts {
        var sum float64
        for j := range accounts[i] {
           sum += float64(accounts[i][j])
        }
        maxNum = math.Max(maxNum, sum)
    }
    return int(maxNum)
}