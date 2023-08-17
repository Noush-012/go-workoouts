package integer

import "fmt"

func SplitIntegerDigitMath(num int) {
	// num := 1234567
	var res []int
	for num > 0 {
		dig := num % 10

		res = append(res, dig)
		num /= 10
	}
	j := len(res) - 1
	for i := 0; i < (len(res)-1)/2; i++ {
		res[i], res[j] = res[j], res[i]
		j--
	}
	fmt.Println(res)
}

func FibonacciR(n int) int {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return FibonacciR(n-1) + FibonacciR(n-2)

}
