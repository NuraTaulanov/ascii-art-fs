package Output

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func OpenFile(s string) ([]string, error) {
	file := fmt.Sprintf("./fonts/%s.txt", s)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")

	return lines, nil
}
