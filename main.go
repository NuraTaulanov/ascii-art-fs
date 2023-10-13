package main

import (
	"os"
	"strings"
)

const (
	pathStandard = "./fonts/standard.txt"
	hashStandard = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"

	pathShadow = "./fonts/shadow.txt"
	hashShadow = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"

	pathThinkertoy = "./fonts/thinkertoy.txt"
	hashThinkertoy = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		Ascii()

	}

	args = os.Args[1:]
	arguments := strings.Join(args, " ")

	switch {
	case strings.Contains(arguments, "--output"):
		OutputAction()

	case strings.Contains(arguments, "--align"):
		JustifyAction()

	}
}
