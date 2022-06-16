package main

import (
	"errors"
	"fmt"
)

func CalcBin(val1 int, val2 int, fn func(int, int) (int, error)) []int {
	val1Bin, val2Bin := ConvDecToBin(val1), ConvDecToBin(val2)
	val1Len, val2Len := len(val1Bin), len(val2Bin)

	var maxLen int

	if val1Len < val2Len {
		maxLen = val2Len
	} else {
		maxLen = val1Len
	}

	var result []int

	for i := 0; i < maxLen; i++ {

		trgIdx := maxLen - i

		var (
			eval1 int
			eval2 int
		)

		_, _ = eval1, eval2

		if trgIdx > val1Len {
			eval1 = 0
		} else {
			eval1 = val1Bin[i+val1Len-maxLen]
		}

		if trgIdx > val2Len {
			eval2 = 0
		} else {
			eval2 = val2Bin[i+val2Len-maxLen]
		}

		part, err := fn(eval1, eval2)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		result = append(result, part)
	}
	return result
}

func And(val1 int, val2 int) (int, error) {
	if val1 != 0 && val1 != 1 && val2 != 0 && val2 != 1 {
		return -1, errors.New("value must be 0 or 1")
	} else {
		if val1 == 1 && val2 == 1 {
			return 1, nil
		} else {
			return 0, nil
		}
	}
}

func Or(val1 int, val2 int) (int, error) {
	if val1 != 0 && val1 != 1 && val2 != 0 && val2 != 1 {
		return -1, errors.New("value must be 0 or 1")
	} else {
		if val1 == 1 || val2 == 1 {
			return 1, nil
		} else {
			return 0, nil
		}
	}
}

func Xor(val1 int, val2 int) (int, error) {
	if val1 != 0 && val1 != 1 && val2 != 0 && val2 != 1 {
		return -1, errors.New("value must be 0 or 1")
	} else {
		if (val1 == 1 && val2 == 0) || (val1 == 0 && val2 == 1) {
			return 1, nil
		} else {
			return 0, nil
		}
	}
}

func Mod100(els []int) int {
	sum := 0
	for _, el := range els {
		sum += el
	}
	return sum % 100
}
