func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    num, res, rem := x, 0, 0
    for x > 0 {
        rem = x % 10
        res = (res * 10) + rem
        x = x / 10
    }
    return num == res
}