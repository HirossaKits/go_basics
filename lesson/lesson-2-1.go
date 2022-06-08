package lesson

func convDecToBin(val int) []int {
	var result []int
	num := val
	for {
		quotient := num / 2
		mod := num % 2

		result = append([]int{mod}, result[0:]...)

		num = quotient
		if quotient == 0 || quotient == 1 {
			break
		}
	}
	return result
}

func main() {
	convDecToBin(100)
}