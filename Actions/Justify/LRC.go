package Justify

import (
	"fmt"

	"strings"
)

func RightLeftCenter(text, font, align string) {
	termWidth := GetTermWidth()
	res := ""

	args := strings.Split(text, "\\n")

	for _, word := range args {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				res += GetLine(1+int(letter-' ')*9+i, font)
			}
			if align == "left" {
				fmt.Println(res)
			}
			if align == "right" {
				fmt.Print(PrintSpaces(termWidth - len(res)))
				fmt.Println(res)
			}
			if align == "center" || align == "centre" {
				fmt.Print(PrintSpaces(termWidth/2 - (len(res) / 2)))
				fmt.Print(res)
				fmt.Print(PrintSpaces(termWidth/2 - (len(res) / 2)))
				fmt.Println()
			}
			res = ""
		}
	}
}
