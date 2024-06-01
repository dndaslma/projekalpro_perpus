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

type Date struct {
	tanggal, bulan, tahun int
}

func DuluanAtauSama(pinjam, kembali Date) bool {
	if pinjam.tahun > kembali.tahun {
		return false
	} else if pinjam.tahun < kembali.tahun {
		return true
	}

	if pinjam.bulan > kembali.bulan {
		return false
	} else if pinjam.bulan < kembali.bulan {
		return true
	}

	if pinjam.tanggal > kembali.tanggal {
		return false
	}

	return true
}
