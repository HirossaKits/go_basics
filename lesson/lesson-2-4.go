package main

import (
	"errors"
	"fmt"
)

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

func CountSumLessThan(n int, s int) int {
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

func sumSlice(elems []int) int {
	sum := 0
	for _, e := range elems {
		sum += e
	}
	return sum
}

var (
	match  int
	tSum   int
	elem   []int
	result [][]int
)

func searchAll(pre []int) {

	fmt.Printf("%v", pre)

	idx := len(pre)

	if len(elem) < idx {
		if sumSlice(pre) == tSum {
			result = append(result, pre)
		}
		return
	}

	pre1 := append(pre, elem[idx])
	searchAll(pre1)

	pre2 := append(pre, 0)
	searchAll(pre2)

}

func CountSumIs(n int, s int, e []int) {

	match = n
	tSum = s
	elem = e

	searchAll(a := []int{})
}
