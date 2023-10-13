package main

import (
	"ascii/Actions/Fs"
	"ascii/Actions/Justify"
	"ascii/Actions/Output"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func OutputAction() {
	str := os.Args[1]
	fileName := str[9:]
	text := os.Args[2]
	banner := os.Args[3]
	a, err := Output.Output(text, banner)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fileName, []byte(a), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func JustifyAction() {
	font := "standard.txt"
	align := "left"
	text := ""
	if !strings.Contains(os.Args[1], "--align=") {
		fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\nExample: go run . --align=right  something  standard")
		return
	}
	for _, v := range os.Args[1:] {
		if len(v) >= 8 && v[0:8] == "--align=" {
			align = v[8:]
		} else if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			text += v
		}
	}
	text = strings.ReplaceAll(text, "\n", `\n`)
	if align == "left" || align == "right" || align == "center" || align == "centre" {
		Justify.RightLeftCenter(text, font, align)
	} else if align == "justify" {
		lines := strings.Split(text, "\\n")
		for _, v := range lines {
			Justify.Justify(v, font)
		}

	}
}

func Ascii() {
	result := Fs.GetArt(os.Args[1], os.Args[2])
	fmt.Println(result)
}
