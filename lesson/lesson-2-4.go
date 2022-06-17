package main

import "errors"

func CountDivisor(n int, x int, y int) (int, error) {
	if n < 1 || x < 1 || y < 1 {
		return 0, errors.New("invalid parameter arg1, arg2")
	} else {
		dOfx := n / x
		dOfy := n / y
		dOfxy := n / (x * y)

		return dOfx + dOfy - dOfxy, nil
	}
}

func CountLessThan(n int, s int) int {
	count := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i+j <= s {
				count++
			}
		}
	}
	return count
}
