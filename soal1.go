package main

import "fmt"

func main() {
	arr := make([]float64, 0, 9)
	soal1(arr)
}

func soal1(arr []float64) {
	var maxx, minn, avg float64
	arr = append(arr, 5.0, 4.0, 2.5, 5.0, 3.6, 1.1, 3.6, 4.0, 4.2, 1.5)
	total := 0.0
	maxx = arr[0]
	minn = arr[0]

	for i := 0; i < 10; i++ {
		if maxx < arr[i] {
			maxx = arr[i]
		}
		if minn > arr[i] {
			minn = arr[i]
		}
		total += arr[i]
	}
	avg = total / 10
	fmt.Print(minn, " || ", maxx, " || ")
	fmt.Printf("%.2f", avg)
}
