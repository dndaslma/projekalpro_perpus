package utils

import "fmt"

const WrongInputPrompt = "Pilihan anda tidak tersedia, silahkan pilih kembali."
const WaitForEnterPrompt = "\nTekan Enter untuk lanjut."
const NMAX = 5

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func TungguEnter() {
	var s rune
	fmt.Scanf("\n%c", &s)
}
