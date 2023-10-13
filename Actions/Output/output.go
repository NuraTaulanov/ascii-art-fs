package Output

import "errors"

func Output(text, format string) (string, error) {
	flag, err := ValidBanner(format)
	if err != nil {
		return "", err
	}
	if !flag {
		return "", errors.New("someone changed our files")
	}
	template, e := OpenFile(format)
	if e != nil {
		return "", e
	}
	result := asciiart(text, template)
	return result, nil
}
func asciiart(s string, lines []string) string {
	arr := []rune{}
	Newline := false
	result := ""
	for i, r := range s {
		if Newline {
			Newline = false
			result = stringArt(arr, lines, result)
			arr = []rune{}
			continue
		}
		if r == 92 && len(s) != i+1 {
			if s[i+1] == 110 {
				Newline = true
				continue
			}
		}
		arr = append(arr, r)
	}
	return stringArt(arr, lines, result)
}
func stringArt(arr []rune, lines []string, result string) string {
	if len(arr) != 0 {
		for line := 0; line < 8; line++ {
			for _, r := range arr {
				skip := (r-32)*9 + 1
				result = result + lines[line+int(skip)]
			}
			result = result + "\n"
		}
	}
	result = result + "\n"
	return result
}
