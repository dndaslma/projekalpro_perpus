package main

import (
	"fmt"
	"pkg/pengguna"
	"pkg/utils"
)

func main() {
	var angka int
	var p pengguna.ArrIdentitas

	for {
		utils.ClearScreen()
		Menu()
		fmt.Print("Pilih [1/2/3/4]: ")
		fmt.Scan(&angka)

		if angka == 1 {
			pengguna.Menu(&p)
			utils.TungguEnter()
		} else if angka == 2 {
			fmt.Println("Kamu telah memilih nomor 2")
			utils.TungguEnter()
		} else if angka == 3 {
			fmt.Println("Kamu telah memilih nomor 3")
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
	fmt.Println("||     M E N U     ||")
	fmt.Println("=====================")
	fmt.Println("1. Data Pengguna     ")
	fmt.Println("2. Data Buku         ")
	fmt.Println("3. Data Peminjaman   ")
	fmt.Println("4. Keluar            ")
	fmt.Println("=====================")

}
