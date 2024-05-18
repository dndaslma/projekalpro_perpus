package pengguna

import (
	"fmt"
	"pkg/utils"
)

type identitas struct {
	nama, email, nomorHp string
}

type ArrIdentitas [utils.NMAX]identitas

func Menu(A *ArrIdentitas) {
	var data ArrIdentitas
	var nData, angka int
	var nama string
	for {
		utils.ClearScreen()
		printMenu()
		fmt.Print("Pilih [1/2/3/4/5]: ")
		fmt.Scan(&angka)
		if angka == 1 {
			bacaPengguna(&data, &nData)
			fmt.Print(utils.WaitForEnterPrompt)
			utils.TungguEnter()
		} else if angka == 2 {
			fmt.Print("Masukkan data yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusData(&data, &nData, nama)
			fmt.Print(utils.WaitForEnterPrompt)
			utils.TungguEnter()
		} else if angka == 3 {
			fmt.Print("Masukkan data yang ingin diubah: ")
			fmt.Scan(&nama)
			ubahData(&data, nData, nama)
			fmt.Print(utils.WaitForEnterPrompt)
			utils.TungguEnter()
		} else if angka == 4 {
			cetakPengguna(data, nData)
			fmt.Print(utils.WaitForEnterPrompt)
			utils.TungguEnter()
		} else if angka == 5 {
			break
		} else {
			fmt.Println(utils.WrongInputPrompt)
		}
	}

}

func bacaPengguna(A *ArrIdentitas, n *int) {
	var jumlah int
	fmt.Println("Masukkan jumlah pengguna yang ingin ditambahkan:")
	fmt.Scan(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&A[*n].nama)
		fmt.Print("Masukkan Email: ")
		fmt.Scan(&A[*n].email)
		fmt.Print("Masukkan Nomor HP: ")
		fmt.Scan(&A[*n].nomorHp)
		*n++
		fmt.Println("masukkan data selanjutnya") //cara supaya ini waktu for terakhir ga ikut
	}
}

func cetakPengguna(A ArrIdentitas, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Nomor HP: ", A[i].nomorHp)
		fmt.Println()
	}
}

func hapusData(A *ArrIdentitas, n *int, X string) {
	var d int
	d = caridata(*A, *n, X)
	if d != -1 {
		for i := d; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}

}

func caridata(A ArrIdentitas, n int, X string) int {
	var found int = -1
	var med int
	var low int = 0
	var high int = n - 1
	for low <= high && found == -1 {
		med = (low + high) / 2
		if X < A[med].nama {
			high = med - 1
		} else if X > A[med].nama {
			low = med + 1
		} else {
			found = med
		}
	}
	return found
}

func ubahData(A *ArrIdentitas, n int, X string) {
	var d int
	d = caridata(*A, n, X)
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

func printMenu() {
	fmt.Println("====================")
	fmt.Println("||    M E N U     ||")
	fmt.Println("====================")
	fmt.Println("1. Tambah Data      ")
	fmt.Println("2. Hapus Data       ")
	fmt.Println("3. Ubah Data        ")
	fmt.Println("4. Tampilkan Data   ")
	fmt.Println("5. Kembali          ")
	fmt.Println("====================")
}
