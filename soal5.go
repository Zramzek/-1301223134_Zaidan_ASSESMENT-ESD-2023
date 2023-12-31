package main

import (
	"fmt"
	"regexp"
)

func main() {
	var uname string

	fmt.Scan(&uname)
	if len(uname) < 6 {
		uname = bersihinString(uname)
		fmt.Print(soal5(len(uname)))
	}
}

func bersihinString(uname string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(uname, "")
}

func soal5(panjang int) int {
	if panjang == 1 {
		return panjang
	}
	return panjang * soal5(panjang-1)
}
