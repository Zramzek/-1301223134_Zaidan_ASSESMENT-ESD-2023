package main

import (
	"fmt"
	"strings"
)

func main() {
	bool1 := soal2("Aku Suka Kamu")
	bool2 := soal2("Ibu Ratna antar ubi")
	if !bool1 {
		fmt.Println("suka blyat")
	} else {
		fmt.Println("eureeka!")
	}

	if !bool2 {
		fmt.Println("suka blyat")
	} else {
		fmt.Println("eureeka!")
	}
}

func soal2(kalimat string) bool {
	kalimat = strings.ToLower(kalimat)
	pKalimat := len(kalimat) - 1

	for i := 0; i < pKalimat/2; i++ {
		if kalimat[i] != kalimat[pKalimat] {
			return false
		}
		pKalimat--
	}
	return true
}
