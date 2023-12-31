package main

import "fmt"

func main() {
	arr := make([]int, 0, 20)
	arr = append(arr, 20, 1, 3, 2, 4, 6, 8, 5, 7, 9, 11, 13, 15, 10, 12, 14, 16, 18, 20, 17, 19)

	fmt.Print(soal4(arr))
}

func soal4(arr []int) bool {
	pArr := len(arr)

	for i := 0; i < pArr; i++ {
		for j := 0; j < pArr; j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}
