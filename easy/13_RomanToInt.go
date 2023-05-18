package easy

func romanToInt(s string) int {
    if s == "" {
        return 0
    }
    value := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    var sum int
    prevNumeral := s[len(s)-1]
    numeral := prevNumeral
    for i := len(s)-1; i >= 0; i-- {
        prevNumeral = numeral
        numeral = s[i]
        if value[numeral] < value[prevNumeral] {
            sum -= value[numeral]
            continue
        }
        sum += value[numeral]
    }
    return sum
}