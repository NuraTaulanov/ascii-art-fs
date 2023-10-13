package Justify

import (
	"fmt"

	"strings"
)

func Justify(words string, font string) {
	termWidth := GetTermWidth()
	words = strings.TrimSpace(words)
	words = RemoveExtraSpaces(words)
	sws := SplitWhiteSpaces(words)
	ar := make([][]string, len(sws))
	j := 0
	res := ""
	for i := 0; i < 8; i++ {
		j = 0
		for _, letter := range words {
			if letter != ' ' {
				res += GetLine(1+int(letter-' ')*9+i, font)
			} else if letter == ' ' && res != "" {
				ar[j] = append(ar[j], res)
				res = ""
				j++
			}
		}
		ar[j] = append(ar[j], res)
		res = ""
	}
	textLen := 0
	for _, arOfStr := range ar {
		textLen += len(arOfStr[0])
	}
	if len(sws) == 1 {
		RightLeftCenter(words, font, "left")
		return
	}
	numSpaces := (termWidth - textLen) / (len(sws) - 1)
	for i := 0; i < 8; i++ {
		for k, v := range ar {
			fmt.Print(v[i])
			if k != len(ar)-1 {
				fmt.Print(PrintSpaces(numSpaces))
			}
		}
		fmt.Println()
	}
}
func RemoveExtraSpaces(input string) string {

	words := strings.Fields(input)

	result := strings.Join(words, " ")

	return result
}
