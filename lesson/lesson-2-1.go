package main

import (
	"math"
)

// Convert Decimal to Binary
func ConvDecToBin(val int) []int {
	var result []int
	num := val
	for num > 0 {
		mod := num % 2
		result = append([]int{mod}, result[0:]...)
		num = num / 2
	}
	return result
}

// Convert Binary to Decimal
func ConvBinToDec(val []int) int {
	var sum int
	for i := 0; i < len(val); i++ {
		sum += int(math.Pow(2, float64(len(val)-1-i)) * float64(val[i]))
	}
	return sum
}
