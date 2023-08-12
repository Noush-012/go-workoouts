package stringOps

import (
	"fmt"
	"strconv"
	"strings"
)

func MakeFirstLetterCapsUsingSlice(s string) {
	word := strings.Fields(s)
	var res string

	for i, v := range word {
		if v[0] >= 'a' && v[0] <= 'z' {
			res = string(rune(v[0] - 32))
			word[i] = res + v[1:]

		} else {
			word[i] = v
		}
	}
	// fmt.Println("----", word)
	// fmt.Println("77", strings.Join(word, " "))
}

func MakeFirstLetterCaps(s string) string {

	var builder strings.Builder

	word := strings.Fields(s)
	fmt.Println(word)
	var res string
	for _, v := range word {

		builder.WriteString(strings.Title(v + " "))
	}
	res = builder.String()

	return res
}

func BinaryAddition(a, b string) string {

	aInt, _ := strconv.ParseInt(a, 2, 0)
	bInt, _ := strconv.ParseInt(b, 2, 0)

	sum := aInt + bInt

	res := strconv.FormatInt(sum, 2)
	fmt.Println(res)
	return res

}

func BinaryAdditionMath(a, b string) string {
	// Pad the shorter binary string with leading zeros
	maxLength := len(a)
	if len(b) > maxLength {
		maxLength = len(b)
	}
	a = padWithZeros(a, maxLength)
	b = padWithZeros(b, maxLength)

	// Initialize variables for carrying and result
	carry := 0
	result := ""

	// Perform binary addition
	for i := maxLength - 1; i >= 0; i-- {
		bitA := int(a[i] - '0')
		bitB := int(b[i] - '0')

		sum := bitA + bitB + carry
		result = string(sum%2+'0') + result
		carry = sum / 2
	}

	// Add final carry if necessary
	if carry > 0 {
		result = "1" + result
	}

	return result

}
func padWithZeros(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}

func IsPalindrome(s string) bool {
	return false
}
