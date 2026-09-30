package luhn

import (
    "fmt"
    "strings"
    "unicode"
)

// Reverse returns the provided string reversed rune-by-rune.
// This properly supports multi-byte Unicode characters (like accented letters or emojis).
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Valid(id string) bool {
	id_t := strings.Join(strings.Fields(id), "")
    
    if len(id_t) <= 1 {
        return false // id must be longer than one character
    }
	
    for _, c := range id_t {
		if !unicode.IsDigit(c) {
			return false // all id chars must be numeric
		}
    }

    id_l := ""
	for i, c := range Reverse(id_t) {
        num := int(c - '0')
        if i % 2 == 1 {
            num = num * 2
            if num > 9 {
                num = num - 9
            }
        }
        id_l = fmt.Sprintf("%s%d", id_l, num)
    }

    sum := 0
	for _, c := range id_l {
        num := int(c - '0') 
        sum += num
	}
    
    if sum % 10 == 0 {
        return true
    }

    return false
}
