package main

import "fmt"

type Lapangan struct {
	ID      int
	Nama    string
	Harga   int
	Status  string
	Penyewa string
}

type Penyewaan struct {
	ID           int
	NamaLapangan string
	NamaPenyewa  string
	JamMulai     string
	JamSelesai   string
	Tanggal      string
	TotalHarga   int
	Status       string
}

func scanString() string {
	var input string
	fmt.Scanln(&input)
	return input
}
func main() {

	var lapangan [10]Lapangan
	var jumlahLapangan = 0

	var penyewaan [50]Penyewaan
	var jumlahPenyewaan = 0

	lapangan[0] = Lapangan{ID: 1, Nama: "Lapangan A", Harga: 100000, Status: "Tersedia", Penyewa: "-"}
	lapangan[1] = Lapangan{ID: 2, Nama: "Lapangan B", Harga: 120000, Status: "Tersedia", Penyewa: "-"}
	lapangan[2] = Lapangan{ID: 3, Nama: "Lapangan C", Harga: 150000, Status: "Tersedia", Penyewa: "-"}
	jumlahLapangan = 3

	var pilihan int

	for {
		fmt.Println("\n========== FUTSAL-BOOK ==========")
		fmt.Println("1. Lihat Daftar Lapangan")
		fmt.Println("2. Tambah Lapangan")
		fmt.Println("3. Edit Lapangan")
		fmt.Println("4. Hapus Lapangan")
		fmt.Println("5. Sewa Lapangan")
		fmt.Println("6. Lihat Riwayat Penyewaan")
		fmt.Println("7. Edit Penyewaan")
		fmt.Println("8. Hapus Penyewaan")
		fmt.Println("9. Cari Penyewa (Binary Search)")
		fmt.Println("10. Urutkan Lapangan by Harga (Selection Sort)")
		fmt.Println("11. Lihat Statistik")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu (1-12): ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			lihatLapangan(lapangan, jumlahLapangan)

		case 2:
			tambahLapangan(&lapangan, &jumlahLapangan)

		case 3:
			editLapangan(&lapangan, &jumlahLapangan)

		case 4:
			hapusLapangan(&lapangan, &jumlahLapangan)

		case 5:
			if jumlahPenyewaan < 50 {
				fmt.Print("Masukkan ID Lapangan yang akan disewa: ")
				var idLapangan int
				fmt.Scan(&idLapangan)

				var indexLapangan = -1
				for i := 0; i < jumlahLapangan; i++ {
					if lapangan[i].ID == idLapangan {
						indexLapangan = i

					}
				}

				if indexLapangan == -1 {
					fmt.Println("Lapangan tidak ditemukan!")
				} else if lapangan[indexLapangan].Status == "Booked" {
					fmt.Println("Lapangan sedang disewa!")
				} else {
					fmt.Print("Masukkan Nama Penyewa: ")
					var namaPenyewa string
					fmt.Scan(&namaPenyewa)

					fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
					var tanggal string
					fmt.Scan(&tanggal)

					fmt.Print("Masukkan Jam Mulai (HH:MM): ")
					var jamMulai string
					fmt.Scan(&jamMulai)

					fmt.Print("Masukkan Jam Selesai (HH:MM): ")
					var jamSelesai string
					fmt.Scan(&jamSelesai)

					totalHarga := lapangan[indexLapangan].Harga * 1

					penyewaan[jumlahPenyewaan] = Penyewaan{
						ID:           jumlahPenyewaan + 1,
						NamaLapangan: lapangan[indexLapangan].Nama,
						NamaPenyewa:  namaPenyewa,
						JamMulai:     jamMulai,
						JamSelesai:   jamSelesai,
						Tanggal:      tanggal,
						TotalHarga:   totalHarga,
						Status:       "Aktif",
					}

					lapangan[indexLapangan].Status = "Booked"
					lapangan[indexLapangan].Penyewa = namaPenyewa

					jumlahPenyewaan++
					fmt.Println("✓ Penyewaan berhasil dicatat!")
					fmt.Printf("Total Harga: Rp %d\n", totalHarga)
				}
			} else {
				fmt.Println("Kapasitas penyewaan penuh!")
			}

		case 6:
			lihatPenyewaan(penyewaan, jumlahPenyewaan)

		case 7:
			editPenyewaan(&penyewaan, jumlahPenyewaan)

		case 8:
			hapusPenyewaan(&penyewaan, &jumlahPenyewaan)

		case 9:
			fmt.Print("Masukkan nama penyewa yang dicari: ")
			var cariNama string
			fmt.Scan(&cariNama)

			// Binary Search (array harus diurutkan dulu)
			urutkanPenyewaByNama(&penyewaan, jumlahPenyewaan)
			index := binarySearchPenyewa(penyewaan, jumlahPenyewaan, cariNama)

			if index != -1 {
				fmt.Printf("\nPenyewa ditemukan!\n")
				fmt.Printf("ID Transaksi: %d\n", penyewaan[index].ID)
				fmt.Printf("Nama: %s\n", penyewaan[index].NamaPenyewa)
				fmt.Printf("Lapangan: %s\n", penyewaan[index].NamaLapangan)
				fmt.Printf("Tanggal: %s\n", penyewaan[index].Tanggal)
				fmt.Printf("Jam: %s - %s\n", penyewaan[index].JamMulai, penyewaan[index].JamSelesai)
			} else {
				fmt.Println("X Penyewa tidak ditemukan!")
			}

		case 10:
			fmt.Println("\nMengurutkan lapangan berdasarkan harga (Selection Sort)...")
			selectionSortLapangan(&lapangan, jumlahLapangan)
			fmt.Println("✓ Lapangan berhasil diurutkan!")
			lihatLapangan(lapangan, jumlahLapangan)

		case 11:
			hitungStatistik(lapangan, jumlahLapangan, penyewaan, jumlahPenyewaan)

		case 12:
			fmt.Println("Terima kasih telah menggunakan Futsal-Book!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahLapangan(lapangan *[10]Lapangan, jumlahLapangan *int) {
	if *jumlahLapangan >= 10 {
		fmt.Println("Kapasitas lapangan penuh!")
		return
	}

	fmt.Print("Masukkan ID Lapangan: ")
	var id int
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Lapangan: ")
	nama := scanString()

	fmt.Print("Masukkan Harga (per jam): ")
	var harga int
	fmt.Scan(&harga)

	lapangan[*jumlahLapangan] = Lapangan{
		ID:      id,
		Nama:    nama,
		Harga:   harga,
		Status:  "Tersedia",
		Penyewa: "-",
	}
	*jumlahLapangan++
	fmt.Println("✓ Lapangan berhasil ditambahkan!")
}

func editLapangan(lapangan *[10]Lapangan, jumlahLapangan *int) {
	lihatLapangan(*lapangan, *jumlahLapangan)
	fmt.Print("\nMasukkan ID Lapangan yang akan diedit: ")
	var idEdit int
	fmt.Scan(&idEdit)

	var indexEdit = -1
	for i := 0; i < *jumlahLapangan; i++ {
		if lapangan[i].ID == idEdit {
			indexEdit = i

		}
	}

	if indexEdit == -1 {
		fmt.Println("Lapangan tidak ditemukan!")

	} else {
		fmt.Println("\nPilih yang akan diedit:")
		fmt.Println("1. Edit Nama")
		fmt.Println("2. Edit Harga")
		fmt.Println("3. Edit Status")
		fmt.Print("Pilih (1-3): ")
		var pilihEdit int
		fmt.Scan(&pilihEdit)

		switch pilihEdit {
		case 1:
			fmt.Print("Masukkan Nama Lapangan baru: ")
			fmt.Scan(&lapangan[indexEdit].Nama)
			fmt.Println("✓ Nama berhasil diubah!")

		case 2:
			fmt.Print("Masukkan Harga baru: ")
			fmt.Scan(&lapangan[indexEdit].Harga)
			fmt.Println("✓ Harga berhasil diubah!")

		case 3:
			fmt.Print("Masukkan Status baru (Tersedia/Booked): ")
			fmt.Scan(&lapangan[indexEdit].Status)
			fmt.Println("✓ Status berhasil diubah!")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

	}
}

func hapusLapangan(lapangan *[10]Lapangan, jumlahLapangan *int) {
	lihatLapangan(*lapangan, *jumlahLapangan)

	fmt.Print("\nMasukkan ID Lapangan yang akan dihapus: ")
	var idHapus int
	fmt.Scan(&idHapus)

	var indexHapus = -1

	for i := 0; i < *jumlahLapangan && indexHapus == -1; i++ {
		if lapangan[i].ID == idHapus {
			indexHapus = i
		}
	}

	// 2. Menggunakan else sebagai pengganti return
	if indexHapus == -1 {
		fmt.Println("Lapangan tidak ditemukan!")
	} else {
		for i := indexHapus; i < *jumlahLapangan-1; i++ {
			lapangan[i] = lapangan[i+1]
		}
		*jumlahLapangan--
		fmt.Println("✓ Lapangan berhasil dihapus!")
	}
}

func editPenyewaan(penyewaan *[50]Penyewaan, jumlah int) {
	lihatPenyewaan(*penyewaan, jumlah)

	fmt.Print("\nMasukkan ID Transaksi yang akan diedit: ")
	var idEdit int
	fmt.Scan(&idEdit)

	var indexEdit = -1
	for i := 0; i < jumlah; i++ {
		if penyewaan[i].ID == idEdit {
			indexEdit = i

		}
	}

	if indexEdit == -1 {
		fmt.Println("Transaksi tidak ditemukan!")

	} else {
		fmt.Println("\nPilih yang akan diedit:")
		fmt.Println("1. Edit Nama Penyewa")
		fmt.Println("2. Edit Tanggal")
		fmt.Println("3. Edit Jam Mulai")
		fmt.Println("4. Edit Jam Selesai")
		fmt.Println("5. Edit Status")
		fmt.Print("Pilih (1-5): ")
		var pilihEdit int
		fmt.Scan(&pilihEdit)

		switch pilihEdit {
		case 1:
			fmt.Print("Masukkan Nama Penyewa baru: ")
			fmt.Scan(&penyewaan[indexEdit].NamaPenyewa)
			fmt.Println("✓ Nama Penyewa berhasil diubah!")

		case 2:
			fmt.Print("Masukkan Tanggal baru (DD-MM-YYYY): ")
			fmt.Scan(&penyewaan[indexEdit].Tanggal)
			fmt.Println("✓ Tanggal berhasil diubah!")

		case 3:
			fmt.Print("Masukkan Jam Mulai baru (HH:MM): ")
			fmt.Scan(&penyewaan[indexEdit].JamMulai)
			fmt.Println("✓ Jam Mulai berhasil diubah!")

		case 4:
			fmt.Print("Masukkan Jam Selesai baru (HH:MM): ")
			fmt.Scan(&penyewaan[indexEdit].JamSelesai)
			fmt.Println("✓ Jam Selesai berhasil diubah!")

		case 5:
			fmt.Print("Masukkan Status baru (Aktif/Selesai): ")
			fmt.Scan(&penyewaan[indexEdit].Status)
			fmt.Println("✓ Status berhasil diubah!")

		default:
			fmt.Println("Pilihan tidak valid!")
		}

	}

}

func hapusPenyewaan(penyewaan *[50]Penyewaan, jumlahPenyewaan *int) {
	lihatPenyewaan(*penyewaan, *jumlahPenyewaan)

	fmt.Print("\nMasukkan ID Transaksi yang akan dihapus: ")
	var idHapus int
	fmt.Scan(&idHapus)

	var indexHapus = -1
	for i := 0; i < *jumlahPenyewaan; i++ {
		if penyewaan[i].ID == idHapus {
			indexHapus = i
			break
		}
	}

	if indexHapus == -1 {
		fmt.Println("Transaksi tidak ditemukan!")
		return
	}

	for i := indexHapus; i < *jumlahPenyewaan-1; i++ {
		penyewaan[i] = penyewaan[i+1]
	}
	*jumlahPenyewaan--

	fmt.Println("✓ Transaksi berhasil dihapus!")
}

func lihatLapangan(lapangan [10]Lapangan, jumlah int) {
	fmt.Println("\n========== DAFTAR LAPANGAN ==========")
	if jumlah == 0 {
		fmt.Println("Belum ada lapangan!")
		return
	}

	fmt.Printf("%-4s %-15s %-12s %-12s %-20s\n", "ID", "Nama", "Harga", "Status", "Penyewa")
	fmt.Println("--------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-4d %-15s Rp%-10d %-12s %-20s\n",
			lapangan[i].ID,
			lapangan[i].Nama,
			lapangan[i].Harga,
			lapangan[i].Status,
			lapangan[i].Penyewa)
	}
}

func lihatPenyewaan(penyewaan [50]Penyewaan, jumlah int) {
	fmt.Println("\n========== RIWAYAT PENYEWAAN ==========")
	if jumlah == 0 {
		fmt.Println("Belum ada riwayat penyewaan!")
		return
	}

	fmt.Printf("%-4s %-15s %-15s %-12s %-12s %-15s %-10s\n",
		"ID", "Lapangan", "Penyewa", "Tanggal", "Jam", "Total Harga", "Status")
	fmt.Println("----------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-4d %-15s %-15s %-12s %s-%s Rp%-9d %-10s\n",
			penyewaan[i].ID,
			penyewaan[i].NamaLapangan,
			penyewaan[i].NamaPenyewa,
			penyewaan[i].Tanggal,
			penyewaan[i].JamMulai,
			penyewaan[i].JamSelesai,
			penyewaan[i].TotalHarga,
			penyewaan[i].Status)
	}
}

func selectionSortLapangan(lapangan *[10]Lapangan, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlah; j++ {
			if lapangan[j].Harga < lapangan[minIndex].Harga {
				minIndex = j
			}
		}
		temp := lapangan[i]
		lapangan[i] = lapangan[minIndex]
		lapangan[minIndex] = temp
	}
}

func urutkanPenyewaByNama(penyewaan *[50]Penyewaan, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlah; j++ {
			if penyewaan[j].NamaPenyewa < penyewaan[minIndex].NamaPenyewa {
				minIndex = j
			}
		}
		temp := penyewaan[i]
		penyewaan[i] = penyewaan[minIndex]
		penyewaan[minIndex] = temp
	}
}

func binarySearchPenyewa(penyewaan [50]Penyewaan, jumlah int, nama string) int {
	left := 0
	right := jumlah - 1

	for left <= right {
		mid := (left + right) / 2

		if penyewaan[mid].NamaPenyewa == nama {
			return mid
		} else if penyewaan[mid].NamaPenyewa < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func hitungStatistik(lapangan [10]Lapangan, jumlahLapangan int,
	penyewaan [50]Penyewaan, jumlahPenyewaan int) {
	fmt.Println("\n========== STATISTIK ==========")

	fmt.Printf("Total Lapangan: %d\n", jumlahLapangan)

	var tersedia = 0
	for i := 0; i < jumlahLapangan; i++ {
		if lapangan[i].Status == "Tersedia" {
			tersedia++
		}
	}
	fmt.Printf("Lapangan Tersedia: %d\n", tersedia)
	fmt.Printf("Lapangan Booked: %d\n", jumlahLapangan-tersedia)

	fmt.Printf("\nTotal Transaksi: %d\n", jumlahPenyewaan)

	var totalPendapatan = 0
	for i := 0; i < jumlahPenyewaan; i++ {
		totalPendapatan += penyewaan[i].TotalHarga
	}
	fmt.Printf("Total Pendapatan: Rp %d\n", totalPendapatan)

	if jumlahLapangan > 0 {
		var totalHarga = 0
		for i := 0; i < jumlahLapangan; i++ {
			totalHarga += lapangan[i].Harga
		}
		rataRata := totalHarga / jumlahLapangan
		fmt.Printf("Rata-rata Harga Lapangan: Rp %d\n", rataRata)
	}
}
