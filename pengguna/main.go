package pengguna

import (
	"fmt"
	"pkg/utils"
)

type identitas struct {
	nama, email, nomorHp string
}

type ArrPengguna [utils.NMAX]identitas

func Utama(A *ArrPengguna, n *int) {
	var angka int
	var nama string
	for {
		utils.ClearScreen()
		PrintMenu()
		fmt.Print("Pilih [1/2/3/4/5]: ")
		fmt.Scan(&angka)
		if angka == 1 {
			BacaPengguna(A, n)
		} else if angka == 2 {
			fmt.Print("Masukkan data yang ingin dihapus (nama): ")
			fmt.Scan(&nama)
			HapusPengguna(A, n, nama)
		} else if angka == 3 {
			fmt.Print("Masukkan data yang ingin diubah (nama): ")
			fmt.Scan(&nama)
			UbahPengguna(A, *n, nama)
		} else if angka == 4 {
			CetakSeluruhPengguna(*A, *n)
		} else if angka == 5 {
			break
		} else {
			fmt.Println(utils.WrongInputPrompt)
		}
		fmt.Print(utils.WaitForEnterPrompt)
		utils.TungguEnter()
	}
}

func BacaPengguna(A *ArrPengguna, n *int) {
	var jumlah int
	var masuk identitas
	fmt.Print("Masukkan jumlah pengguna yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&masuk.nama)

		if CariPenggunaDariNama(*A, *n, masuk.nama) != -1 {
			fmt.Println("Mohon maaf data sudah ada.")
			utils.TungguEnter()
			continue
		}

		fmt.Print("Masukkan Email: ")
		fmt.Scan(&masuk.email)
		fmt.Print("Masukkan Nomor HP: ")
		fmt.Scan(&masuk.nomorHp)

		A[*n] = masuk

		*n++
		if i < jumlah-1 {
			fmt.Println("masukkan data selanjutnya")
		}
	}
}

func CetakSeluruhPengguna(A ArrPengguna, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Nomor HP: ", A[i].nomorHp)
		fmt.Println()
	}
}

func HapusPengguna(A *ArrPengguna, n *int, nama string) {
	var d = CariPenggunaDariNama(*A, *n, nama)
	if d != -1 {
		for i := d; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Pengguna berhasil dihapus.")
	} else {
		fmt.Println("Pengguna tidak ditemukan.")
	}

}

func CariPenggunaDariNama(A ArrPengguna, n int, nama string) int {
	var found int = -1
	var med int
	var low int = 0
	var high int = n - 1
	for low <= high && found == -1 {
		med = (low + high) / 2
		if nama < A[med].nama {
			high = med - 1
		} else if nama > A[med].nama {
			low = med + 1
		} else {
			found = med
		}
	}
	return found
}

func UbahPengguna(A *ArrPengguna, n int, nama string) {
	var d = CariPenggunaDariNama(*A, n, nama)
	if d != -1 {
		fmt.Print("Masukkan Nama baru: ")
		fmt.Scan(&A[d].nama)
		fmt.Print("Masukkan Email baru: ")
		fmt.Scan(&A[d].email)
		fmt.Print("Masukkan Nomor HP baru: ")
		fmt.Scan(&A[d].nomorHp)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func PenggunaAda(P *ArrPengguna, n int, namaPengguna string) bool {
	for i := 0; i < n; i++ {
		if P[i].nama == namaPengguna {
			return true
		}
	}
	return false
}

func PrintMenu() {
	fmt.Println("====================")
	fmt.Println("||    PENGGUNA    ||")
	fmt.Println("====================")
	fmt.Println("1. Tambah Data      ")
	fmt.Println("2. Hapus Data       ")
	fmt.Println("3. Ubah Data        ")
	fmt.Println("4. Tampilkan Data   ")
	fmt.Println("5. Kembali          ")
	fmt.Println("====================")
}
