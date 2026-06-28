func groupAnagrams(strs []string) [][]string {
    var resMap = make(map[[26]int][]string)

    for _, str := range strs {
        var charCount [26]int
        for _, c := range str {
            charCount[c-'a']++
        }
        resMap[charCount] = append(resMap[charCount], str)
    }
    
    var res = make([][]string, 0)
    for _, v := range resMap {
        res = append(res, v)
    }

    return res
}
