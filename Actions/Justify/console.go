package Justify

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	// heigth, err := strconv.Atoi(sArr[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}
