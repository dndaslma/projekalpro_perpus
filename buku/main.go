package buku

import (
	"fmt"
	"pkg/utils"
)

type dBuku struct {
	judul, kodeBuku, kategori string
}

type ArrBuku [utils.NMAX]dBuku

func Utama(A *ArrBuku, n *int) {
	var angka int
	var buku, kode string
	for {
		utils.ClearScreen()
		PrintMenu()
		fmt.Print("Pilih [1/2/3/4/5]: ")
		fmt.Scan(&angka)
		if angka == 1 {
			BacaBuku(A, n)
		} else if angka == 2 {
			fmt.Print("Masukkan buku yang ingin dihapus (judul): ")
			fmt.Scan(&buku)
			HapusBuku(A, n, buku)
		} else if angka == 3 {
			fmt.Print("MasukkanA yang ingin dicari (kode): ")
			fmt.Scan(&kode)
			CetakBuku(*A, *n, kode)
		} else if angka == 4 {
			CetakSeluruhBuku(*A, *n)
		} else if angka == 5 {
			break
		} else {
			fmt.Println(utils.WrongInputPrompt)
		}
		fmt.Print(utils.WaitForEnterPrompt)
		utils.TungguEnter()
	}
}

func BacaBuku(A *ArrBuku, n *int) {
	var jumlah, i int
	var data dBuku
	fmt.Print("Masukkan jumlah buku yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	for i < jumlah {
		fmt.Print("Masukkan Judul: ")
		fmt.Scan(&data.judul)

		if CariBukuDariJudul(*A, *n, data.judul) != -1 {
			fmt.Println("Mohon maaf judul sudah ada.")
			utils.TungguEnter()
			continue
		}

		if CariBukuDariKode(*A, *n, data.kodeBuku) != -1 {
			fmt.Println("Mohon maaf kode sudah ada")
			utils.TungguEnter()
			continue
		}

		fmt.Print("Masukkan Kode Buku: ")
		fmt.Scan(&data.kodeBuku)
		fmt.Print("Masukkan Kategori: ")
		fmt.Scan(&data.kategori)

		A[*n] = data
		*n++
		if i < jumlah-1 {
			fmt.Println("Masukkan buku selanjutnya")
		}
		i++
	}
}

func BukuAda(P *ArrBuku, n int, kode string) bool {
	for i := 0; i < n; i++ {
		if P[i].kodeBuku == kode {
			return true
		}
	}
	return false
}

func CetakSeluruhBuku(A ArrBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Judul: ", A[i].judul)
		fmt.Println("Kode Buku: ", A[i].kodeBuku)
		fmt.Println("Kategori: ", A[i].kategori)
		fmt.Println()
	}
}

func CetakBuku(A ArrBuku, n int, kode string) {
	var index = CariBukuDariKode(A, n, kode)
	if index != -1 {
		fmt.Println("Judul: ", A[index].judul)
		fmt.Println("Kode Buku: ", A[index].kodeBuku)
		fmt.Println("Kategori: ", A[index].kategori)
		fmt.Println()
	} else {
		fmt.Println("Buku Tidak Ditemukan.")
	}
}

func HapusBuku(A *ArrBuku, n *int, judul string) {
	var d = CariBukuDariJudul(*A, *n, judul)
	if d != -1 {
		for i := d; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Buku berhasil dihapus.")
	} else {
		fmt.Println("Buku tidak ditemukan.")
	}

}

func CariBukuDariJudul(A ArrBuku, n int, judul string) int {
	var found int = -1
	var med int
	var low int = 0
	var high int = n - 1
	for low <= high && found == -1 {
		med = (low + high) / 2
		if judul < A[med].judul {
			high = med - 1
		} else if judul > A[med].judul {
			low = med + 1
		} else {
			found = med
		}
	}
	return found
}

func CariBukuDariKode(A ArrBuku, n int, kode string) int {
	var found int = -1
	var med int
	var low int = 0
	var high int = n - 1
	for low <= high && found == -1 {
		med = (low + high) / 2
		if kode < A[med].kodeBuku {
			high = med - 1
		} else if kode > A[med].kodeBuku {
			low = med + 1
		} else {
			found = med
		}
	}
	return found
}

func PrintMenu() {
	fmt.Println("====================")
	fmt.Println("||    B U K U     ||")
	fmt.Println("====================")
	fmt.Println("1. Tambah Buku      ")
	fmt.Println("2. Hapus Buku       ")
	fmt.Println("3. Cari Buku        ")
	fmt.Println("4. Tampilkan Buku   ")
	fmt.Println("5. Kembali          ")
	fmt.Println("====================")
}
