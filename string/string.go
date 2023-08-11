package stringOps

import (
	"fmt"
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
