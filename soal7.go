package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(soal7("xfqfr bfmdz"))
	fmt.Println(soal7("gjxtp lzj rfz ifkyfw jxi snm"))
	fmt.Println(soal7("gwt, gjxtp qz rfz rfpfs in bfwlty lfp?"))
	fmt.Println(soal7("fpz xfdfsl pfrz, rfz lfp ofin ufhfwpz"))
	fmt.Println(soal7("dfsl pnwnr xynhpjw otrtp pz pnhp ifwn lwzu"))
}

func soal7(kalimat string) string {
	var decryptedMessage strings.Builder

	for _, char := range kalimat {
		if char >= 'a' && char <= 'z' {
			decryptedChar := 'a' + (char-'a'-5+26)%26
			decryptedMessage.WriteRune(decryptedChar)
		} else if char >= 'A' && char <= 'Z' {
			decryptedChar := 'A' + (char-'A'-5+26)%26
			decryptedMessage.WriteRune(decryptedChar)
		} else {
			decryptedMessage.WriteRune(char)
		}
	}

	return decryptedMessage.String()
}
