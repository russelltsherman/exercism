package phonenumber

import (
    "errors"
    "fmt"
    "regexp"
)

func Number(phoneNumber string) (string, error) {
    // Compile regex matching everything except digits 0-9
	reg := regexp.MustCompile("[^0-9]+")
	res := reg.ReplaceAllString(phoneNumber, "")

    if len(res) < 10 {
        return res, errors.New("number too short")
    }
    
    if len(res) > 10 {
        if res[0:1] == "1" {
            res = res[1:] // trim us country code
        } else {
	        return res, errors.New("number too long")        
        }
    }

    if res[0:1] == "0" || res[0:1] == "1" {
        return res, errors.New("area code invalid")
    }

    if res[3:4] == "0" || res[3:4] == "1" {
        return res, errors.New("exchange code invalid")
    }
    
	return res, nil
}

func AreaCode(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
	
    res := fmt.Sprintf("%s", num[0:3])
    return res, err
}

func Format(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)

    res := fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:])
    return res, err
}
