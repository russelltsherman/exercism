package isbnverifier

import (
    "strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
    
    length := len(isbn)
    if length != 10 {
        return false // isbn should be 10 digits
    }

    sum := 0
    for _, num := range isbn {
        val := int(num - '0')
        if val == 40 { // rune is X
            if length > 1 { 
                return false // rune is not final character
            }
            val = 10
        }
        sum = sum + val * length

		length --
    } 
    
    return sum % 11 == 0
}
