func canConstruct(ransomNote string, magazine string) bool {
    c := make([]int, 26)

    for _, s := range magazine {
        c[s - 97]++
    }
       
    for _, s := range ransomNote {
        if c[s - 97] == 0 {
            return false
        }
        c[s - 97]--
    }

    return true
}