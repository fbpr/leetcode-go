func romanToInt(s string) int {
    roman := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    total, temp := 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if roman[s[i]] < temp {
            total -= roman[s[i]]
        } else {
            total += roman[s[i]]
        }
        temp = roman[s[i]]
    }
    return total
}