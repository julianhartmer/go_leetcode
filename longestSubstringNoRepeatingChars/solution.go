func rune2idx(c rune) int {
    return int(c) - int('a');
}

func maxInt(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func lengthOfLongestSubstring(s string) int {
    begin := 0
    result := 0

    for p,c := range s {
        for i := p - 1; i >= begin; i-- {
            if rune(s[i]) == c {
                result = maxInt(result, p - begin)
                begin = i + 1
            }
        }
    }
    return maxInt(result, len(s) - begin)
}
