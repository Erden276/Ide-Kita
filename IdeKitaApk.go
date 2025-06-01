package main

import "fmt"

type proyek struct {
	id       int
	nama     string
	status   string
	judul    string
	kategori string
	upvote   int
	tanggal  string
}

type array [1000]proyek

func main() {
	var p array
	var n, lanjut int
	var no int = 0
	var pilih int
	fmt.Println("Hallo Selamat Datang di Aplikasi kami")
	fmt.Println("Silahkan Pilih Ide Yang Kamu Sarankan")
	for n <= 3 {
		data(&n)
		switch n {
		case 1:
			tambahdata(&p, no)
			no++
		case 2:
			if no == 0 {
				fmt.Println("Data Tidak Ada")
				clear()
			} else {
				tampilkandata(p, no)
				fmt.Println("Silahkan Pilih: ")
				fmt.Println("1. Sorting")
				fmt.Println("2. Hapus")
				fmt.Println("3. Cari")
				fmt.Println("4. Edit")
				fmt.Println("5. Voting")
				fmt.Println("6. Tampilkan Data Berdasarkan ID")
				fmt.Println("7. Tampilkan Ide Populer Berdasarkan Periode")
				fmt.Println("8. Back")
				fmt.Print("Silahkan Pilih Opsi: ")
				fmt.Scan(&lanjut)
				switch lanjut {
				case 1:
					selectionsort(&p, no)
					tampilkandata(p, no)
				case 2:
					hapus(&p, &no)
				case 3:
					fmt.Println("Cari Sesuai: ")
					fmt.Println("1. Nama")
					fmt.Println("2. Kategori")
					fmt.Println("3. ID")
					fmt.Println("4. Back")
					fmt.Print("Pilih: ")
					fmt.Scan(&pilih)
					if pilih <= 2 {
						linearsearch(&p, no, pilih)
					} else if pilih == 3 {
						insertionsort(&p, no)
						binarysearch(&p, no)
					} else {
						break
					}
				case 4:
					edit(&p, no)
					tampilkandata(p, no)
				case 5:
					upvote(&p, no)
				case 6:
					var idCari int
					fmt.Print("Masukkan ID yang ingin ditampilkan: ")
					fmt.Scan(&idCari)
					tampilkanDetailByID(p, no, idCari)
				case 7:
					tampilkanPopulerPeriode(p, no)
				case 8:
					clear()
					break
				}
			}
		default:
			return
		}
	}
}

func tampilkanPopulerPeriode(p array, no int) {
	var tAwal, tAkhir string
	fmt.Print("Masukkan tanggal awal (YYYY-MM-DD): ")
	fmt.Scan(&tAwal)
	fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ")
	fmt.Scan(&tAkhir)

	var hasil array
	var count, i, j int
	for i = 0; i < no; i++ {
		if p[i].tanggal >= tAwal && p[i].tanggal <= tAkhir {
			hasil[count] = p[i]
			count++
		}
	}
	for i = 0; i < count; i++ {
		for j = i + 1; j < count; j++ {
			if hasil[i].upvote < hasil[j].upvote {
				hasil[i], hasil[j] = hasil[j], hasil[i]
			}
		}
	}
	tampilkandata(hasil, count)
}

func data(n *int) {
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Tampilkan Data")
	fmt.Println("3. Exit")
	fmt.Print("Silahkan Pilih Opsi: ")
	fmt.Scan(&*n)
}
func tampilkanDetailByID(p array, no int, id int) {
	var found bool = false
	var i int
	for i = 0; i < no; i++ {
		if p[i].id == id {
			fmt.Println()
			fmt.Println("ID      :", p[i].id)
			fmt.Println("Judul   :", p[i].judul)
			fmt.Println("Nama    :", p[i].nama)
			fmt.Println("Status  :", p[i].status)
			fmt.Println("Kategori:", p[i].kategori)
			fmt.Println("Upvote  :", p[i].upvote)
			fmt.Println("Tanggal :", p[i].tanggal)
			fmt.Println()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
	}
}

func tambahdata(p *array, no int) {
	var pilih, idBaru int
	var n int
	var m error
	for {
		fmt.Print("ID: ")
		n, m = fmt.Scan(&idBaru)
		if m != nil || n != 1 {
			fmt.Println("ID harus berupa angka, tidak boleh huruf!")
			var inIdAgain string
			fmt.Scan(&inIdAgain)
			continue
		}
		if cariIndexByID(p, no, idBaru) != -1 {
			fmt.Println("ID sudah digunakan, silakan masukkan ID lain!")
		} else {
			p[no].id = idBaru
			break
		}
	}
	fmt.Print("Judul: ")
	fmt.Scan(&p[no].judul)
	fmt.Print("Nama: ")
	fmt.Scan(&p[no].nama)
	fmt.Print("Status: ")
	fmt.Scan(&p[no].status)
	fmt.Println("Pilih Kategori: ")
	fmt.Println("1. Produk")
	fmt.Println("2. Pemasaran")
	fmt.Println("3. Operasional")
	fmt.Println("4. Teknologi")
	fmt.Println("5. Lainnya")
	fmt.Println("6. Back")
	fmt.Print("Silahkan Pilih Kategori: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		p[no].kategori = "Produk"
	case 2:
		p[no].kategori = "Pemasaran"
	case 3:
		p[no].kategori = "Operasional"
	case 4:
		p[no].kategori = "Teknologi"
	case 5:
		fmt.Print("Masukkan nama kategori baru: ")
		fmt.Scan(&p[no].kategori)
	default:
		return
	}
	clear()
	fmt.Print("Tanggal dibuat (YYYY-MM-DD): ")
	fmt.Scan(&p[no].tanggal)
	for {
		fmt.Println("Apa yang ingin Anda lakukan selanjutnya?")
		fmt.Println("1. Upvote Ide")
		fmt.Println("2. Lanjut")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			upvote(p, no+1)
		case 2:
			clear()
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func upvote(p *array, jumlah int) {
	var id, idx int
	tampilkandata(*p, jumlah)
	fmt.Print("Masukkan ID ide yang ingin di-upvote: ")
	fmt.Scan(&id)
	idx = cariIndexByID(p, jumlah, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	p[idx].upvote++
	fmt.Println("Upvote berhasil!")
}

func tampilkandata(p array, no int) {
	clear()
	colom()
	var i int
	for i = 0; i < no; i++ {
		fmt.Printf("| %-2d | %-5d | %-25s | %-15s | %-15s | %-20s | %-6d | %-10s |\n", i+1, p[i].id, p[i].judul, p[i].nama, p[i].status, p[i].kategori, p[i].upvote, p[i].tanggal)
	}
	fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
}

func cariIndexByID(p *array, jumlah int, id int) int {
	var i int
	for i = 0; i < jumlah; i++ {
		if p[i].id == id {
			return i
		}
	}
	return -1
}

func insertionsort(A *array, n int) {
	var i, j int
	var temp proyek
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].id > temp.id {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func selectionsort(A *array, no int) {
	var i, j, sc int
	clear()
	fmt.Println("Pilih Sorting: ")
	fmt.Println("1. Upvote Tertinggi")
	fmt.Println("2. Upvote Terendah")
	fmt.Println("3. Tanggal Terbaru")
	fmt.Println("4. Tanggal Terlama")
	fmt.Print("Silahkan Pilih Opsi: ")
	fmt.Scan(&sc)
	switch sc {
	case 1:
		for i = 0; i < no; i++ {
			for j = i + 1; j < no; j++ {
				if A[i].upvote < A[j].upvote {
					A[i], A[j] = A[j], A[i]
				}
			}
		}
	case 2:
		for i = 0; i < no; i++ {
			for j = i + 1; j < no; j++ {
				if A[i].upvote > A[j].upvote {
					A[i], A[j] = A[j], A[i]
				}
			}
		}
	case 3:
		for i = 0; i < no; i++ {
			for j = i + 1; j < no; j++ {
				if A[i].tanggal < A[j].tanggal {
					A[i], A[j] = A[j], A[i]
				}
			}
		}
	case 4:
		for i = 0; i < no; i++ {
			for j = i + 1; j < no; j++ {
				if A[i].tanggal > A[j].tanggal {
					A[i], A[j] = A[j], A[i]
				}
			}
		}
	default:
		fmt.Println("Pilihan sorting tidak valid.")
	}
}

func colom() {
	fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
	fmt.Printf("| %-2s | %-5s | %-25s | %-15s | %-15s | %-20s | %-6s | %-10s |\n", "NO", "ID", "Judul", "Nama", "Status", "Kategori", "Upvote", "Tanggal")
	fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
}

func linearsearch(p *array, no int, pilih int) {
	var nama, kategori string
	var i int
	var found bool
	switch pilih {
	case 1:
		fmt.Print("Cari Nama: ")
		fmt.Scan(&nama)
		found = false
		for i = 0; i < no; i++ {
			if p[i].nama == nama {
				colom()
				fmt.Printf("| %-2d | %-5d | %-25s | %-15s | %-15s | %-20s | %-6d | %-10s |\n", i+1, p[i].id, p[i].judul, p[i].nama, p[i].status, p[i].kategori, p[i].upvote, p[i].tanggal)
				found = true
			}
		}
		fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
		if !found {
			fmt.Println("Data tidak ditemukan.")
		}
	case 2:
		fmt.Print("Cari Kategori: ")
		fmt.Scan(&kategori)
		found = false
		for i = 0; i < no; i++ {
			if p[i].kategori == kategori {
				colom()
				fmt.Printf("| %-2d | %-5d | %-25s | %-15s | %-15s | %-20s | %-6d | %-10s |\n", i+1, p[i].id, p[i].judul, p[i].nama, p[i].status, p[i].kategori, p[i].upvote, p[i].tanggal)
				found = true
			}
		}
		fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
		if !found {
			fmt.Println("Data tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func binarysearch(p *array, n int) {
	var left, mid, right, id int
	var found bool = false
	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scan(&id)
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if p[mid].id == id {
			if p[mid].id == id {
				clear()
				found = true
				colom()
				fmt.Printf("| %-2d | %-5d | %-25s | %-15s | %-15s | %-20s | %-6d | %-10s |\n", mid+1, p[mid].id, p[mid].judul, p[mid].nama, p[mid].status, p[mid].kategori, p[mid].upvote, p[mid].tanggal)
				fmt.Println("|-------------------------------------------------------------------------------------------------------------------------|")
				left = right + 1
			} else if p[mid].id > id {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if !found {
			fmt.Println("Data Tidak Ditemukan")
		}
	}
}

func hapus(p *array, no *int) {
	clear()
	var id, idx, i int
	var konfirmasi string
	tampilkandata(*p, *no)
	fmt.Print("Masukkan ID data yang ingin dihapus: ")
	fmt.Scan(&id)
	idx = cariIndexByID(p, *no, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	fmt.Println("Detail ide yang akan dihapus:")
	fmt.Println("ID      :", p[idx].id)
	fmt.Println("Judul   :", p[idx].judul)
	fmt.Println("Nama    :", p[idx].nama)
	fmt.Println("Status  :", p[idx].status)
	fmt.Println("Kategori:", p[idx].kategori)
	fmt.Println("Upvote  :", p[idx].upvote)
	fmt.Println("Tanggal :", p[idx].tanggal)

	fmt.Print("Yakin ide ini mau dihapus? (y/n): ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" || konfirmasi == "Y" {
		for i = idx; i < *no-1; i++ {
			p[i] = p[i+1]
		}
		*no--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}
func clear() {
	fmt.Print("\033[H\033[2J")
}

func edit(p *array, no int) {
	var pilih, id, idx, kat int
	if no == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa diedit.")
		return
	}
	tampilkandata(*p, no)
	fmt.Print("Masukkan ID ide yang ingin diedit: ")
	fmt.Scan(&id)
	idx = cariIndexByID(p, no, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}
	for {
		fmt.Println("Edit Bagian Mana: ")
		fmt.Println("1. Judul")
		fmt.Println("2. Nama")
		fmt.Println("3. Status")
		fmt.Println("4. Kategori")
		fmt.Println("5. Back")
		fmt.Print("Silahkan Pilih: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Print("Silahkan Ganti Judul: ")
			fmt.Scan(&p[idx].judul)
		case 2:
			fmt.Print("Silahkan Ganti Nama: ")
			fmt.Scan(&p[idx].nama)
		case 3:
			fmt.Print("Silahkan Ganti Status: ")
			fmt.Scan(&p[idx].status)
		case 4:
			fmt.Println("Pilih Kategori: ")
			fmt.Println("1. Produk")
			fmt.Println("2. Pemasaran")
			fmt.Println("3. Operasional")
			fmt.Println("4. Teknologi")
			fmt.Println("5. Lainnya")
			fmt.Print("Silahkan Ganti: ")
			fmt.Scan(&kat)
			switch kat {
			case 1:
				p[idx].kategori = "Produk"
			case 2:
				p[idx].kategori = "Pemasaran"
			case 3:
				p[idx].kategori = "Operasional"
			case 4:
				p[idx].kategori = "Teknologi"
			case 5:
				fmt.Print("Masukkan nama kategori baru: ")
				fmt.Scan(&p[idx].kategori)
			default:
				fmt.Println("Pilihan kategori tidak valid.")
			}
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
