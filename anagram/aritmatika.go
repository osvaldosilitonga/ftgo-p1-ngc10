package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	var choice string

	// Input
	fmt.Println("a. Penjumlahan")
	fmt.Println("b. Pengurangan")
	fmt.Println("c. Perkalian")
	fmt.Println("d. Pembagian")
	// Input Pilih Operasi
	fmt.Print("Pilih operasi aritmatika : ")
	fmt.Scan(&choice)

	switch choice {
	case "a":
		angka := inputUser()
		result := Penjumlahan(angka[0], angka[1])
		fmt.Printf("Hasil : %.2f", result)
	case "b":
		angka := inputUser()
		result := Pengurangan(angka[0], angka[1])
		fmt.Printf("Hasil : %.2f", result)
	case "c":
		angka := inputUser()
		result := Perkalian(angka[0], angka[1])
		fmt.Printf("Hasil : %.2f", result)
	case "d":
		angka := inputUser()
		result := Pembagian(angka[0], angka[1])
		fmt.Printf("Hasil : %.2f", result)
	default:
		panic("Operasi aritmatika tidak ditemukan!")
	}

}

func inputUser() []float64 {
	var nilai1, nilai2 float64

	// Input Angka
	fmt.Print("Angka 1 : ")
	_, errNilai1 := fmt.Scan(&nilai1)
	if errNilai1 != nil {
		panic("Inputan harus angka!")
	}

	fmt.Print("Angka 2 : ")
	_, errNilai2 := fmt.Scan(&nilai2)
	if errNilai2 != nil {
		panic("Inputan harus angka!")
	}
	return []float64{nilai1, nilai2}
}

func Penjumlahan(n1, n2 float64) float64 {
	return n1 + n2
}
func Pengurangan(n1, n2 float64) float64 {
	return n1 - n2
}
func Perkalian(n1, n2 float64) float64 {
	return n1 * n2
}
func Pembagian(n1, n2 float64) float64 {
	return n1 / n2
}
