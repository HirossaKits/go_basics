package main

import (
	"fmt"
	"math"
)

// Convert Decimal to Binary
func convDecToBin(val int) []int {
	var result []int
	num := val
	for {
		quotient := num / 2
		mod := num % 2
		result = append([]int{mod}, result[0:]...)

		num = quotient
		if quotient == 0 || quotient == 1 {
			result = append([]int{quotient}, result[0:]...)
			break
		}
	}
	return result
}

// Convert Binary to Decimal
func convBinToDec(val []int) int {
	var sum int
	for i := 0; i < len(val) -1; i++ {
		sum += int(math.Pow(2,float64(len(val) - 1 - i)) * float64(val[i]))
	}
	return sum
}


func main() {
	res := convDecToBin(100)
	fmt.Printf("%d\n",res)

	res2 := convBinToDec(res)
	fmt.Printf("%d\n",res2)
}