func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }
    
    // sort.Slice(strs, func(i, j int) bool {
    //     return len(strs[i]) < len(strs[j])
    // })

    // lcp, comp := "", strs[0]
    // for i := 0; i < len(comp); i++ {
    //     for j := 1; j < len(strs); j++ {
    //         if strs[j][i] != comp[i] {
    //            return lcp
    //         } 
    //     }

    //     lcp += string(comp[i])
    // }

    // strings built-in
    lcp := strs[0]
    for i := 0; i < len(strs); i++{
        for !strings.HasPrefix(strs[i], lcp){
            lcp = lcp[:len(lcp) - 1]
        }
    }

    return lcp
}
