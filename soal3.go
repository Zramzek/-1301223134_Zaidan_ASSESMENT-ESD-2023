package main

import "fmt"

func main() {
	tamu := []string{"Ningguang", "Hutao", "Xiao", "Childe"}
	kebiasaan := map[string]string{
		"Ningguang": "Memeriksa kue",
		"Hutao":     "Memberikan kado tanpa memperhatikan kue",
		"Xiao":      "Memotret apa pun yang dilihat pertama kali",
		"Childe":    "Membawa air mineral dan meletakkannya di meja",
	}
	soal3(tamu, kebiasaan)
}

func soal3(tamu []string, kebiasaan map[string]string) {
	fotoXiao := "Kue ada di meja."
	var tersangka string

	for _, tamu := range tamu {
		kebiasaanTamu := kebiasaan[tamu]

		if tamu == "Xiao" && fotoXiao == "Kue masih utuh di meja." {
			continue
		}

		if kebiasaanTamu == "Memeriksa kue" {
			tersangka = tamu
			break
		}
	}
	fmt.Printf("Berdasarkan analisis, %s adalah yang paling mungkin mengambil kue.\n", tersangka)
}
