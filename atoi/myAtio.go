func myAtoi(s string) int {
    i := 0
    value := 0
    isPositive := true

    for i = 0; i < len(s) && s[i] == ' '; i++ {
    }

    if i < len(s) && s[i] == '+'  {
        i += 1
    } else if i < len(s) && s[i] == '-' {
        isPositive = false
        i += 1
    }

    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        value = value * 10 + (int) (s[i] - '0')
        if isPositive && math.MaxInt32 < value {
            return math.MaxInt32
        }
        if !isPositive && math.MinInt32 > -value {
            return math.MinInt32
        }
        i++
    }

    if isPositive {
        return value
    } else {
        return -value
    }
}
