package main

import (
	"fmt"
)

type Barang struct {
	produk, kategori string
	harga            int
}

type Pelanggan struct {
	nama     string
	arrMinat []string
	arrBeli  []string
}

var arrBarang []Barang
var arrPelanggan []Pelanggan

func main() {
	fmt.Println("Rina : ", rekomendasi("Rina"))
	fmt.Println("Budi : ", rekomendasi("Budi"))
	fmt.Println("Hartono : ", rekomendasi("Hartono"))
}

func rekomendasi(nama string) []string {
	arrBarang = []Barang{
		{"TV", "elektronik", 1000},
		{"headphone", "elektronik", 200},
		{"baju", "fashion", 50},
		{"gitar", "musik", 300},
		{"sepatu", "olahraga", 80},
		{"kamera", "elektronik", 600},
	}

	arrPelanggan = []Pelanggan{
		{"Rina", []string{"elektronik", "musik"}, []string{"TV", "headphone"}},
		{"Budi", []string{"fashion", "musik"}, []string{"baju", "gitar"}},
		{"Hartono", []string{"olahraga", "elektronik"}, []string{"sepatu", "kamera"}},
	}

	arrRekomen := make([]string, 0)

	for _, rangePelanggan := range arrPelanggan {
		if nama == rangePelanggan.nama {
			for _, rangeKategori := range rangePelanggan.arrMinat {
				for _, rangeMinat := range arrBarang {
					if rangeKategori == rangeMinat.kategori {
						arrRekomen = append(arrRekomen, rangeMinat.produk)
					}
				}
			}
		}
	}
	return arrRekomen
}
