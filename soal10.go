package main

import (
	"fmt"
)

func main() {
	totalPembayaran1 := 10000
	totalBelanja1 := 7500
	fmt.Println(hitungKembalian(totalPembayaran1, totalBelanja1))

	totalPembayaran2 := 5000
	totalBelanja2 := 1100
	fmt.Println(hitungKembalian(totalPembayaran2, totalBelanja2))

	totalPembayaran3 := 178000
	totalBelanja3 := 90500
	fmt.Println(hitungKembalian(totalPembayaran3, totalBelanja3))
}

func hitungKembalian(totalPembayaran, totalBelanja int) map[string]int {
	pecahanUang := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := totalPembayaran - totalBelanja
	hasil := make(map[string]int)

	for _, pecahan := range pecahanUang {
		jumlahPecahan := kembalian / pecahan
		if jumlahPecahan > 0 {
			hasil[fmt.Sprintf("%d", pecahan)] = jumlahPecahan
			kembalian %= pecahan
		}
	}

	return hasil
}
