package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    i := 0
    if n < 1 {
        return 0, errors.New("input less than 1") 
    }
    for n > 1 {
    	i ++
        n = next(n)
    }
    return i, nil
}

func next(n int) int {
	if n % 2 == 0 {
        n = n / 2
    } else {
        n = (n * 3) + 1
    }
    return n
}
