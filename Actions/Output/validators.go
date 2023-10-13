package Output

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

const (
	hash_shadow     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	hash_standard   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	hash_thinkertoy = "6153b98a3b781e765b86286a7e8c8bc8aa52fc636dcdc81c56ccc0866db58af3"
)

func ValidString(input string) bool {
	for _, ch := range input {
		if (int(ch) < 32 || int(ch) > 126) && int(ch) != 10 && int(ch) != 9 {
			return false
		}
	}
	return true
}

func ValidBanner(format string) (bool, error) {
	hashStr, err := hash(format)
	if err != nil {
		return false, err
	}
	if hashStr == hash_shadow || hashStr == hash_standard || hashStr == hash_thinkertoy {
		return true, nil
	}
	return false, nil
}

func hash(format string) (string, error) {
	file := fmt.Sprintf("./fonts/%s.txt", format)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(content)

	hashString := fmt.Sprintf("%x", hash)
	return hashString, nil
}
