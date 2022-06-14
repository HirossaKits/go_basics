package main

import (
	"fmt"
)

func main() {
	res1 := ConvDecToBin(101)
	fmt.Printf("%d\n",res1)

	res2 := ConvBinToDec(res1)
	fmt.Printf("%d\n",res2)

	res3 := CalcAnd(13,14)
	fmt.Println(res3)

	res4 := ConvBinToDec(res3)
	fmt.Println(res4)
}