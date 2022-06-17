package main

import (
	"fmt"
)

func main() {
	res1 := ConvDecToBin(101)
	fmt.Printf("%d\n", res1)

	res2 := ConvBinToDec(res1)
	fmt.Printf("%d\n", res2)

	res3 := ConvBinToDec(CalcBin(13, 14, And))
	fmt.Println(res3)

	res4 := ConvBinToDec(CalcBin(13, 14, Or))
	fmt.Println(res4)

	res5 := ConvBinToDec(CalcBin(13, 14, Xor))
	fmt.Println(res5)

	res6, err := CountDivisor(15, 3, 5)
	if err == nil {
		fmt.Println(res6)
	}

	res7 := CountLessThan(3, 4)
	fmt.Println(res7)

}
