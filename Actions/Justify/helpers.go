package Justify

import (
	"bufio"
	"log"
	"os"
)

func PrintSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func GetLine(num int, filename string) string {
	file, err := os.Open("./fonts/" + filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer file.Close()
	content := bufio.NewScanner(file)
	lineNum := 0
	line := ""
	for content.Scan() {
		if lineNum == num {
			line = content.Text()
		}
		lineNum++
	}
	return line
}
func SplitWhiteSpaces(str string) []string {
	num := 0
	var prevRune rune
	prevRune = ' '
	word := ""
	for _, v := range str {
		if (v != ' ' && v != '\t' && v != '\n') && (prevRune == ' ' || prevRune == '\t' || prevRune == '\n') {
			num++
		}
		prevRune = v
	}
	array := make([]string, num)
	a := 0
	for _, v := range str {
		if v != ' ' && v != '\t' && v != '\n' {
			word = word + string(v)
		}
		if v == ' ' || v == '\t' || v == '\n' && word != "" {
			a++
			array[a-1] = word
			word = ""
		}
	}
	if word != "" {
		array[a] = word
	}
	return array
}
