package main
import "fmt"

const NMAX = 1001
type Datalapangan struct {
	id				int
	nama			string
	harga			int
	statusLapangan	bool
}

type Datapenyewa struct {
	id				int
	nama			string
	noTelepon		string
}

type Jadwalsewa struct {
	id				int
	idlapangan		int
	idPenyewa		int
	tanggal			string
	jamMulai		int
	jamSelesai		int
	harga			int
	statusJadwal	bool
	statusBayar 	bool
}

//tipe array
type ArrLapangan [NMAX]Datalapangan
type ArrPenyewa [NMAX]Datapenyewa
type ArrJadwal [NMAX]Jadwalsewa

var Lapangan Arrlapangan
var nLapangan int

var Penyewa Arrpenyewa
var nPenyewa int

var Jadwal ArrJadwal
var nJadwal int

var idLapanganCounter int
var idPenyewaCounter int
var idJadwalCounter int

//format jamnya
func cetakjam(jam int) {
	fmt.Printf("%02d:%02d", jam/100, jam%100)
}

//A. bagian data lapangan
func MenambahLapangan(nama string, harga int) {
	Lapangan[nLapangan].id = idLapanganCounter
	Lapangan[nLapangan].nama = nama
	Lapangan[nLapangan].harga = harga
	Lapangan[nLapangan].statusLapangan = true
	nLapangan++
	idLapanganCounter++
	fmt.Println("Lapangan berhasil ditambahkan.")
}

func CariIdxLapangan(id int) int {
	var found int
	var i int
	found = -1
	i = 0
	for i < nLapangan && found == -1 {
		if Lapangan[i].id == id {
			found = i
		}
		i = i + 1
	}
	return found
}

func MengubahLapangan(id int, namabaru string, hargabaru int) {
	var idx int 
	idx = CariIdxLapangan(id)
	if idx == -1 {
		fmt.Println("Lapangan tidak ditemukan.")
	} else {
		Lapangan[idx].nama = namabaru
		Lapangan[idx].harga = hargabaru
		fmt.Println("Lapangan berhasil diubah.")
	}
}

func MenghapusLapangan(id int) {
	var idx int
	var i int
	idx = CariIdxLapangan(id)
	if idx == -1 {
		fmt.Println("Lapangan tidak ditemukan.")
	} else {
		i = idx
		for i < nLapangan-1 {
			Lapangan[i] = Lapangan[i+1]
			i = i + 1
		}
		nLapangan--
		fmt.Println("Lapangan berhasil dihapus.")
	}
}

func cetakLapangan() {
	var i int
	var status string
	i =  0
	fmt.Println("=== Data Lapangan ===")
	for i < nLapangan {
		if !Lapangan[i].statusLapangan {
			status = "Tidak Tersedia"
		} else {
			status = "Tersedia"
		}
		fmt.Printf("ID: %d | Nama: %s | Harga: Rp%d/jam | Status: %s\n", Lapangan[i].id, Lapangan[i].nama, Lapangan[i].harga, status)
		i = i + 1
	}
}

//A. bagian data penyewa
func MenambahPenyewa(nama string, noTelepon string) {
	Penyewa[nPenyewa].id = idPenyewaCounter
	Penyewa[nPenyewa].nama = nama
	Penyewa[nPenyewa].noTelepon = noTelepon
	nPenyewa++
	idPenyewaCounter++
	fmt.Println("Penyewa berhasil ditambahkan.")
}

func CariIdxPenyewa(id int) int {
	var found int
	var i int
	found = -1
	i = 0
	for i < nPenyewa && found == -1 {
		if Penyewa[i].id == id {
			found = i
		}
		i = i + 1
	}
	return found
}


func MengubahPenyewa(id int, namabaru,  noTeleponbaru string) {
	var idx int
	idx = CariIdxPenyewa(id)
	if idx == -1 {
		fmt.Println("Penyewa tidak ditemukan.")
	} else {
		Penyewa[idx].nama = namabaru
		Penyewa[idx].noTelepon = noTeleponbaru
		fmt.Println("Penyewa berhasil diubah.") 
	}
}

func MenghapusPenyewa(id int) {
	var idx int
	var i int 
	idx = CariIdxPenyewa(id)
	if idx == -1 {
		fmt.Println("Penyewa tidak ditemukan.")
	} else {
		i = idx
		for i < nPenyewa-1 {
			Penyewa[i] = Penyewa[i+1]
			i = i + 1
		}
		nPenyewa--
		fmt.Println("Penyewa berhasil dihapus.")
	}
}

func cetakPenyewa() {
	var i int
	i = 0
	fmt.Println("=== Data Penyewa ===")
	for i < nPenyewa {
		Fmt.Printf("ID: %d | Nama: %s | No. Telepon: %s\n", Penyewa[i].id, Penyewa[i].nama, Penyewa[i].noTelepon)
		i = i + 1
	}
}

//B. bagian mencatat transaksi penyewaan dan status ketersediaan jam operasional lapangan
func cariIdxJadwal(id int) int {
	var found int
	var i int
	found = -1
	i = 0
	for i < nJadwal && found == -1 {
		if Jadwal[i].id == id {
			found = i 
		}	
		i = i + 1
	}
	return found
}

func TambahJadwal(idLapangan int, idPenyewa int, tanggal string, jamMulai int, jamSelesai int) {
	var idxLapangan int
	var idxPenyewa int
	idxLapangan = CariIdxLapangan(idLapangan)
	idxPenyewa = CariIdxPenyewa(idPenyewa)
	if idxLapangan == -1 {
		fmt.Println("Lapangan tidak ditemukan.")
	} else if idxPenyewa == -1 {
		fmt.Println("Penyewa tidak ditemukan.")
	} else if !Lapangan[idxLapangan].statusLapangan {
		fmt.Println("Lapangan tidak tersedia.")
	} else {
		Jadwal[nJadwal].id = idJadwalCounter
		Jadwal[nJadwal].idLapangan = idLapangan
		Jadwal[nJadwal].idPenyewa = idPenyewa
		Jadwal[nJadwal].tanggal = tanggal
		Jadwal[nJadwal].jamMulai = jamMulai
		Jadwal[nJadwal].jamSelesai = jamSelesai
		Jadwal[nJadwal].harga = Lapangan[idxLapangan].harga
		Jadwal[nJadwal].statusJadwal = true 
		Jadwal[nJadwal].statusBayar = false
		nJadwal++
		idJadwalCounter++
		fmt.Println("Jadwal berhasil ditambahkan.")
	}
}

func BayarJadwal(id int) {
	var idx int
	idx = cariIdxJadwal(id)
	if idx == -1 {
		fmt.Println("Jadwal tidak ditemukan.")
	} else {
		Jadwal[idx].statusBayar = true
		fmt.Println("Pembayaran berhasil.")
	}
}

func cetakJadwal() {
	var i int
	var statusJadwal string
	var statusBayar string 
	i = 0
	fmt.Println("=== Data Jadwal ===")
	for i < nJadwal {
		if Jadwal[i].statusJadwal {
			statusJadwal = "Aktif"
		} else {
			statusJadwal = "Selesai"
		}
		if Jadwal[i].statusBayar {
			statusBayar = "Lunas"
		} else {
			statusBayar = "Belum Bayar"
		}
		fmt.Printf("ID: %d | Lapangan: %d | Penyewa: %d | Tanggal: %s | Jam: ", Jadwal[i].id,  Jadwal[i].idLapangan,  Jadwal[i].idPenyewa,  Jadwal[i].tanggal)
		cetakjam(Jadwal[i].jamMulai)
		fmt.Printf(" - ")
		cetakjam(Jadwal[i].jamSelesai)
		fmt.Printf("| Harga: Rp%d | Status: %s | Bayar: %s\n", Jadwal[i].harga, statusJadwal, statusBayar)
		i = i + 1
	}
}

//C. bagian sequential search berdasarkan Nama & No telpon
func SeqSearchNama(N string) int {
	var found int
	var i int
	found = -1
	i = 0
	for i < nPenyewa && found == -1 {
		if Penyewa[i].nama == N {
			found = i
		}
		i = i + 1
	}	
	return found
}

func SeqSearchNoTelepon(N string) int {
	var found int
	var i int
	found = -1
	i = 0
	for i < nPenyewa && found == -1 {
		if Penyewa[i].noTelepon == N {
			found = i
		}
		i = i + 1
	}	
	return found
}

//C. bagian binary search berdasarkan Nama & No telpon
func BinarySearchNama(N string) int {
	var found, kiri, kanan, tengah int
	found = -1
	kiri = 0
	kanan = nPenyewa - 1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if N < Penyewa[tengah].nama {
			kanan = tengah - 1
		} else if N > Penyewa[tengah].nama {
			kiri = tengah + 1
		} else {
			found = tengah
		}
	}
	return found
}

func BinarySearchNoTelepon(N string) int {
	var found, kiri, kanan, tengah int
	found = -1
	kiri = 0
	kanan = nPenyewa - 1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if N < Penyewa[tengah].noTelepon {
			kanan = tengah - 1
		} else if N > Penyewa[tengah].noTelepon {
			kiri = tengah + 1
		} else {
			found = tengah
		}
	}
	return found
}

//D. bagian selection sort
func SelectionSortjadwallapanganjammulai(ar [NMAX]Jadwalsewa, n int) {
	var i, j, idxmin int
	for i = 0; i < n-1; i++ {
		idxmin = i
		for j = i + 1; j < n; j++ {
			if ar[j].jamMulai < ar[idxmin].jamMulai {
				idxmin = j
			}
		}
		if idxmin != i {
			ar[i], ar[idxmin] = ar[idxmin], ar[i]
		}
	}
}

func selectionsortjadwallapanganharga(ar [NMAX]Jadwalsewa, n int) {
	var i, j, idxmin int
	for i = 0; i < n-1; i++ {
		idxmin = i
		for j = i + 1; j < n; j++ {
			if ar[j].harga < ar[idxmin].harga {
				idxmin = j
			}
		}
		if idxmin != i {
			ar[i], ar[idxmin] = ar[idxmin], ar[i]
		}
	}
}


//D. bagian insertion sort
func InsertionSortjadwallapanganjammulai(ar [NMAX]Jadwalsewa, n int) {
	var i, j int
	var key Jadwalsewa
	for i = 1; i < n; i++ {
		key = ar[i]
		j = i - 1
		for j >= 0 && ar[j].jamMulai > key.jamMulai {
			ar[j+1] = ar[j]
			j = j - 1
		}
		ar[j+1] = key
	}
}

func InsertionSortjadwallapanganharga(ar [NMAX]Jadwalsewa, n int) {

//E. bagian nampilin statistik
func Statistik() {
	
}

func main() {
	var n, i int
	var 
}