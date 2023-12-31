package main

import "fmt"

type Menu struct {
	Nama  string
	Tipe  string
	Harga float64
}

type Pesanan struct {
	Menu
	Jumlah int
}

func main() {
	menu := []Menu{
		{"Ayam Goreng Krispi", "Makanan", 15000},
		{"Ayam Puk Puk (Bukan di Geprek)", "Makanan", 13000},
		{"Ayam Bakar", "Makanan", 20000},
		{"Es Teh", "Minuman", 5000},
		{"Es Jeruk", "Minuman", 7000},
	}

	biayaRehan := hitungBiaya(Pesanan{menu[2], 2}) + hitungBiaya(Pesanan{menu[3], 1})
	biayaAmba := hitungBiaya(Pesanan{menu[1], 1}) + hitungBiaya(Pesanan{menu[3], 3})
	biayaFaiz := hitungBiaya(Pesanan{menu[0], 1}) + hitungBiaya(Pesanan{menu[1], 1}) + hitungBiaya(Pesanan{menu[2], 1}) + hitungBiaya(Pesanan{menu[3], 1}) + hitungBiaya(Pesanan{menu[4], 1})

	totalBiaya := biayaRehan + biayaAmba + biayaFaiz
	totalBiaya += totalBiaya * 0.15

	fmt.Println("Biaya Rehan Whangsap:", biayaRehan)
	fmt.Println("Biaya Amba Roni:", biayaAmba)
	fmt.Println("Biaya Faiz Ngawi:", biayaFaiz)
	fmt.Println("Total Biaya:", int(totalBiaya))
}

func hitungBiaya(pesanan Pesanan) float64 {
	var pajak float64

	switch pesanan.Tipe {
	case "Makanan":
		pajak = 0.05
	case "Minuman":
		pajak = 0.03
	}
	biaya := float64(pesanan.Jumlah) * pesanan.Harga
	biaya += biaya * pajak
	return biaya
}
