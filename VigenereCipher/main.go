package main

import (
	"fmt"
	"strings"
)

var temp string

func cipher(text, keyword string) string {
	message := ""
	keyIndex := 0
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c >= 'A' && c <= 'Z' {

			c -= 'A'
			k := keyword[keyIndex] - 'A'

			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}
		message += string(c)
	}
	return message
}

func decipher(cipherText, keyword string) string {
	message := ""
	keyIndex := 0
	for i := 0; i < len(cipherText); i++ {

		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		message += string(c)

		keyIndex++
		keyIndex %= len(keyword)
	}
	return message
}

func main() {
	var text string
	fmt.Print("hi, let's cipher your text with vigenere method: ")
	fmt.Scanf("%s\n", &text)
	text = strings.ToUpper(text)

	var key string
	fmt.Print("type your keyword: ")
	fmt.Scanf("%s\n", &key)
	key = strings.ToUpper(key)

	temp = cipher(text, key)
	fmt.Println(temp)
	temp = decipher(temp, key)
	fmt.Println(temp)
}
