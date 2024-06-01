package main

import (
	"fmt"
	"pkg/buku"
	"pkg/peminjaman"
	"pkg/pengguna"
	"pkg/utils"
)

func main() {
	var angka int
	var P pengguna.ArrPengguna
	var B buku.ArrBuku
	var M peminjaman.ArrPeminjaman
	var n1, n2, n3 int

	utils.ClearScreen()
	Menu()
	fmt.Print("Pilih [1/2/3/4]: ")
	fmt.Scan(&angka)
	for angka < 4 {
		//utils.ClearScreen()
		//Menu()
		//fmt.Print("Pilih [1/2/3/4]: ")
		//fmt.Scan(&angka)
		utils.ClearScreen()
		Menu()
		fmt.Print("Pilih [1/2/3/4]: ")
		if angka == 1 {
			pengguna.Utama(&P, &n1)
			utils.TungguEnter()
		} else if angka == 2 {
			buku.Utama(&B, &n2)
			utils.TungguEnter()
		} else if angka == 3 {
			peminjaman.Utama(&M, &n3, &P, &n1, &B, &n2)
			utils.TungguEnter()
		} else if angka == 4 {
			break
		} else {
			fmt.Println(utils.WrongInputPrompt)
			utils.TungguEnter()
		}
	}
}

func Menu() {
	fmt.Println("=====================")
	fmt.Println("||   MENU UTAMA    ||")
	fmt.Println("=====================")
	fmt.Println("1. Data Pengguna     ")
	fmt.Println("2. Data Buku         ")
	fmt.Println("3. Data Peminjaman   ")
	fmt.Println("4. Keluar            ")
	fmt.Println("=====================")

}
