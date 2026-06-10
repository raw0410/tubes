package main

import "fmt"

const nmax int = 10

type lapangan struct {
	ID    int
	nama  string
	harga float64
	aktif bool
}

type penyewa struct {
	id      int
	nama    string
	telepon string
}

type jadwal struct {
	id           int
	idLapangan   int
	idPenyewa    int
	tanggal      string
	waktuMulai   string
	waktuSelesai string
	total        float64
}

type arlapangan [nmax]lapangan
type arpenyewa [nmax]penyewa
type arjadwal [nmax]jadwal

func tambahLapangan(l arlapangan, id int, nama string, harga float64, aktif bool) [nmax]lapangan {
	lapanganBaru := lapangan{
		ID:    id,
		nama:  nama,
		harga: harga,
		aktif: aktif,
	}
	for i := 0; i < nmax; i++ {
		if lapangan[i].ID == 0 {
			lapangan[i] = lapanganBaru

		}
	}
	return lapangan
}

func tambahPenyewa(penyewa arpenyewa, id int, nama string, telepon string) [nmax]penyewa {
	penyewaBaru := penyewa{
		id:      id,
		nama:    nama,
		telepon: telepon,
	}
	for i := 0; i < nmax; i++ {
		if penyewa[i].id == 0 {
			penyewa[i] = penyewaBaru
			break
		}
	}
	return penyewa
}

func tambahJadwal(jadwal arjadwal, id int, idLapangan int, idPenyewa int, tanggal string, waktuMulai string, waktuSelesai string, total float64) [nmax]jadwal {
	jadwalBaru := jadwal{
		id:           id,
		idLapangan:   idLapangan,
		idPenyewa:    idPenyewa,
		tanggal:      tanggal,
		waktuMulai:   waktuMulai,
		waktuSelesai: waktuSelesai,
		total:        total,
	}
	for i := 0; i < nmax; i++ {
		if jadwal[i].id == 0 {
			jadwal[i] = jadwalBaru
			break
		}
	}
	return jadwal

}

func catattransaksi(jadwal arjadwal, lapangan arlapangan, penyewa arpenyewa) {
	for i := 0; i < nmax; i++ {
		if jadwal[i].id != 0 {
			var namaLapangan string
			var namaPenyewa string
			for j := 0; j < nmax; j++ {
				if lapangan[j].ID == jadwal[i].idLapangan {
					namaLapangan = lapangan[j].nama
					break
				}
			}

			for k := 0; k < nmax; k++ {
				if penyewa[k].id == jadwal[i].idPenyewa {
					namaPenyewa = penyewa[k].nama
					break
				}
			}
			fmt.Printf("ID Jadwal: %d, Lapangan: %s, Penyewa: %s, Tanggal: %s, Waktu Mulai: %s, Waktu Selesai: %s, Total: %.2f\n",
				jadwal[i].id, namaLapangan, namaPenyewa, jadwal[i].tanggal, jadwal[i].waktuMulai, jadwal[i].waktuSelesai, jadwal[i].total)
		}
	}
}
func statustersedialapangan(lapangan arlapangan) {
	for i := 0; i < nmax; i++ {
		if lapangan[i].ID != 0 {
			status := "Tersedia"
			if !lapangan[i].aktif {
				status = "Tidak Tersedia"
			}

			fmt.Printf("ID: %d, Nama: %s, Harga: %.2f, Status: %s\n", lapangan[i].ID, lapangan[i].nama, lapangan[i].harga, status)
		}
	}
}

func sequensialdatapenyewa(penyewa arpenyewa, id int) {
	for i := 0; i < nmax; i++ {
		if penyewa[i].id == id {
			fmt.Printf("ID: %d, Nama: %s, Telepon: %s\n", penyewa[i].id, penyewa[i].nama, penyewa[i].telepon)
		}
	}
	fmt.Println("Penyewa tidak ditemukan")

}

func binarysearchdatapenyewa(penyewa arpenyewa, id int) {
	kiri := 0
	kanan := nmax - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if penyewa[tengah].id == id {
			fmt.Printf("ID: %d, Nama: %s, Telepon: %s\n", penyewa[tengah].id, penyewa[tengah].nama, penyewa[tengah].telepon)
			return
		} else if penyewa[tengah].id < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("Penyewa tidak ditemukan")
}

func selectionsorturutjadwal(jadwal arjadwal) {
	for i := 0; i < nmax-1; i++ {
		minIndex := i
		for j := i + 1; j < nmax; j++ {
			if jadwal[j].id < jadwal[minIndex].id {
				minIndex = j
			}
		}
		jadwal[i], jadwal[minIndex] = jadwal[minIndex], jadwal[i]
	}
}

func insertionsorturutjadwal(jadwal arjadwal) {
	for i := 1; i < nmax; i++ {
		key := jadwal[i]
		j := i - 1
		for j >= 0 && jadwal[j].id > key.id {
			jadwal[j+1] = jadwal[j]
			j--
		}

		jadwal[j+1] = key
	}
}

func TampilkanStatistik() {
	totalPendapatan := 0

	var jamCounter [24]int

	for i := 0; i < countTransaksi; i++ {
		if dataTransaksi[i].IsSelesai {
			totalPendapatan += dataTransaksi[i].Harga
		}

		jam := dataTransaksi[i].JamSewa
		if jam >= 0 && jam < 24 {
			jamCounter[jam]++
		}
	}

	maxFrekuensi := 0
	jamTerpopuler := -1
	for i := 0; i < 24; i++ {
		if jamCounter[i] > maxFrekuensi {
			maxFrekuensi = jamCounter[i]
			jamTerpopuler = i
		}
	}

	fmt.Printf("Total Pendapatan Bulanan: Rp %d\n", totalPendapatan)
	if jamTerpopuler != -1 {
		fmt.Printf("Jam Paling Sering Dipesan: Jam %02d:00\n", jamTerpopuler)
	} else {
		fmt.Println("Belum ada data transaksi.")
	}
}

func main() {
	var lapangan arlapangan
	var penyewa arpenyewa
	var jadwal arjadwal
	lapangan = tambahLapangan(lapangan, 1, "Lapangan A", 100000, true)
	lapangan = tambahLapangan(lapangan, 2, "Lapangan B", 150000, true)
	penyewa = tambahPenyewa(penyewa, 1, "John Doe", "08123456789")
	penyewa = tambahPenyewa(penyewa, 2, "Jane Smith", "08987654321")
	jadwal = tambahJadwal(jadwal, 1, 1, 1, "2024-06-01", "10:00", "12:00", 200000)
	jadwal = tambahJadwal(jadwal, 2, 2, 2, "2024-06-02", "14:00", "16:00", 300000)
	catattransaksi(jadwal, lapangan, penyewa)
	statustersedialapangan(lapangan)
	sequensialdatapenyewa(penyewa, 1)
	binarysearchdatapenyewa(penyewa, 2)
	selectionsorturutjadwal(jadwal)
	insertionsorturutjadwal(jadwal)
	TampilkanStatistik()
}
