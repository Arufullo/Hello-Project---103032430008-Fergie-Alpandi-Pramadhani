package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 100

type data struct {
	idTransaksi string
	namaBarang  string
	jumlah      int
	totalHarga  int
}
type dataTransaksi [NMAX]data

type dataSimpanPinjam struct {
	idTransaksi string
	namaPelaku string
	status string
	totalTransaksi int
}
type simpanPinjamTransaksi [NMAX]dataSimpanPinjam

// Fungsi utama
func main() {
	var pilihan, jumlahData, jumlahDataSP int
	var A dataTransaksi
	var B simpanPinjamTransaksi
	for pilihan < 3 {
		clearScreen()
		menu()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
			case 1:
				jualBeli(&A, &jumlahData)
			case 2:
				simpanPinjam(&B, &jumlahDataSP)
			case 3:
				hehe()
		}
	}
}

func menu(){
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║                  K O P E R A S I                   ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║      1. Koperasi Jual Beli                         ║")
	fmt.Println("║      2. Koperasi Simpan Pinjam                     ║")
	fmt.Println("║      3. Exit                                       ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Print("Pilih (1/2/3)? ")
}

// Program ini adalah aplikasi sederhana untuk mengelola data transaksi koperasi
func jualBeli(A *dataTransaksi, jumlahData *int) {
	var pilihan, idx int

	for pilihan < 7 {
		clearScreen()
		menuJualBeli()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			baca(A, jumlahData)
		case 2:
			cetakMenu(*A, *jumlahData)
		case 3:
			search(A, jumlahData, &idx)
		case 4:
			editData(A, jumlahData, &idx)
		case 5:
			hapus(A, jumlahData, &idx)
		case 6:
			reset(A, jumlahData)
		}
	}
}

func menuJualBeli() { // Tampilan interaktif
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║          K O P E R A S I  J U A L  B E L I         ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║      1. Input Data Transaksi                       ║")
	fmt.Println("║      2. Cetak Data Transaksi                       ║")
	fmt.Println("║      3. Search                                     ║")
	fmt.Println("║      4. Edit Data                                  ║")
	fmt.Println("║      5. Hapus Data                                 ║")
	fmt.Println("║      6. Reset Data                                 ║")
	fmt.Println("║      7. Exit                                       ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Print("Pilih (1/2/3/4/5/6/7)? ")
}

// Fungsi untuk membersihkan layar
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk menunggu input dari pengguna sebelum melanjutkan
// Ini digunakan untuk memberikan jeda sebelum kembali ke menu utama
func jeda() {
	fmt.Println("Tekan Enter untuk kembali ke menu...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	reader.ReadString('\n')
}

// Fungsi untuk membaca data transaksi dari pengguna
// Data yang dimasukkan adalah nama barang, jumlah, dan total harga
func baca(A *dataTransaksi, jumlahData *int) {
	var namaBarang string
	var jumlah, totalHarga int
	var stop bool = false
	fmt.Println("Silakan masukan 0 0 0 jika ingin berhenti")
	fmt.Println("Nama barang | Jumlah | Total Harga")
	for !stop {
		fmt.Scan(&namaBarang, &jumlah, &totalHarga)
		if namaBarang == "0" && jumlah == 0 && totalHarga == 0 {
			stop = true
		} else if jumlah > 0 || totalHarga > 0 {
			A[*jumlahData].idTransaksi = fmt.Sprintf("KJB%02d", *jumlahData+1) // penginputan id transaksi dilakukan secara otomatis
			A[*jumlahData].namaBarang = namaBarang
			A[*jumlahData].jumlah = jumlah
			A[*jumlahData].totalHarga = totalHarga
			*jumlahData++
		} else {
			fmt.Println("\033[31mJumlah dan total harga tidak boleh kurang dari 0\033[0m")
		}
	}
	fmt.Println("\033[32mData sudah berhasil diinput\033[0m")
	jeda()
	fmt.Println()
}

// Fungsi untuk menampilkan menu cetak
func cetakMenu(A dataTransaksi, jumlahData int) {
	var pilihan int
	for pilihan < 4 {
		clearScreen()
		menuCetak()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			cetakData(A, jumlahData)
		case 2:
			terurutMenurun(A, jumlahData)
		case 3:
			terurutMenaik(A, jumlahData)
		}
	}
}

func menuCetak() {
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║             C E T A K  J U A L  B E L I            ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║      1. Tampilkan Semua Data Transaksi             ║")
	fmt.Println("║      2. Tampilkan Data Transaksi Terurut Menurun   ║")
	fmt.Println("║      3. Tampilkan Data Transaksi Terurut Menaik    ║")
	fmt.Println("║      4. Kembali ke Menu Utama                      ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Print("Pilih (1/2/3/4)? ")
}

// fungsi untuk mencetak header dan footer tabel
func headingTable() {
	fmt.Println("╔══════════════╦══════════════╦════════╦═════════════╗")
	fmt.Printf("║ %-12s ║ %-12s ║ %-6s ║ %-11s ║\n", "ID Transaksi", "Nama Barang", "Jumlah", "Total Harga")
	fmt.Println("╠══════════════╬══════════════╬════════╬═════════════╣")
}

func footerTable() {
	fmt.Println("╚══════════════════════════════════════╩═════════════╝")
}

// fungsi untuk mencetak data transaksi
func cetak(A dataTransaksi, i int) {
	fmt.Printf("║ %-12s ║ %-12s ║ %-6d ║ %-11d ║\n", A[i].idTransaksi, A[i].namaBarang, A[i].jumlah, A[i].totalHarga)
}

// fungsi untuk mencetak semua data transaksi
func cetakData(A dataTransaksi, jumlahData int) {
	clearScreen()
	if A[0] == (data{}) {
		fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
	} else {
		headingTable()
		for i := 0; i < jumlahData; i++ {
			cetak(A, i)
		}
		cetakTotalNilaiTransaksi(A, jumlahData)
		footerTable()
	}
	jeda()
	fmt.Println()
}

// fungsi untuk mencetak total nilai transaksi
// total nilai transaksi dihitung dengan menjumlahkan total harga dari semua data transaksi
func cetakTotalNilaiTransaksi(A dataTransaksi, jumlahData int) {
	fmt.Println("╠══════════════╩══════════════╩════════╬═════════════╣")
	fmt.Printf("║ %-36s ║ %-11d ║\n", "Total Nilai Transaksi", totalNilaiTransaksi(A, jumlahData))
}

// fungsi untuk menghitung total nilai transaksi
// total nilai transaksi dihitung dengan menjumlahkan total harga dari semua data transaksi
func totalNilaiTransaksi(A dataTransaksi, jumlahData int) int {
	var total int
	for i := 0; i < jumlahData; i++ {
		total = total + A[i].totalHarga
	}
	return total
}

func menuKriteriaSort() int {
	var kriteria int
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Jumlah")
	fmt.Println("2. Total Harga")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&kriteria)
	return kriteria
}

// Fungsi untuk mengurutkan data transaksi secara menurun
// Fungsi ini menggunakan insertion sort untuk mengurutkan data berdasarkan kriteria yang dipilih
func terurutMenurun(A dataTransaksi, jumlahData int) {
	clearScreen()
	if jumlahData == 0 {
		fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
		jeda()
	} else {
		var B dataTransaksi
		B = A
		kriteria := menuKriteriaSort()
		insertionSortDesc(&B, jumlahData, kriteria)

		fmt.Println("\033[32mData terurut menurun:\033[0m")
		cetakData(B, jumlahData)
	}
}

func insertionSortDesc(A *dataTransaksi, jumlahData int, kriteria int) {
	for i := 1; i < jumlahData; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && ((kriteria == 1 && A[j].jumlah < key.jumlah) || (kriteria == 2 && A[j].totalHarga < key.totalHarga)) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

// Fungsi untuk mengurutkan data transaksi secara menaik
// Fungsi ini menggunakan selection sort untuk mengurutkan data berdasarkan kriteria yang dipilih
func terurutMenaik(A dataTransaksi, jumlahData int) {
	clearScreen()
	if jumlahData == 0 {
		fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
		jeda()
	} else {
		var B dataTransaksi
		B = A
		kriteria := menuKriteriaSort()
		selectionSort(&B, jumlahData, kriteria)

		fmt.Println("\033[32mData terurut menaik:\033[0m")
		cetakData(B, jumlahData)
	}
}

func selectionSort(A *dataTransaksi, jumlahData int, kriteria int) {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if (kriteria == 1 && A[j].jumlah < A[minIdx].jumlah) ||
				(kriteria == 2 && A[j].totalHarga < A[minIdx].totalHarga) {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

// Fungsi untuk mencari data transaksi berdasarkan ID transaksi atau nama barang
func search(A *dataTransaksi, jumlahData, i *int) {
	var pilihan int

	for pilihan < 3 {
		clearScreen()
		menuSearch()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			cariIdTransaksi(A, jumlahData, i)
			jeda()
		case 2:
			cariNamaBarang(A, jumlahData)
		}
	}
}

// Fungsi untuk menampilkan menu pencarian
func menuSearch() {
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║        P E N C A R I A N  J U A L  B E L I         ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║      1. Berdasarkan ID Transaksi                   ║")
	fmt.Println("║      2. Berdasarkan Nama Barang                    ║")
	fmt.Println("║      3. Exit                                       ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Print("Pilih (1/2/3)? ")
}

// Fungsi untuk mencari data transaksi berdasarkan ID transaksi
// ID transaksi yang dicari dimasukkan oleh pengguna
func cariIdTransaksi(A *dataTransaksi, jumlahData, i *int) {
	var n string

	fmt.Print("Masukan ID Transaksi (contoh: KJB01): ")
	fmt.Scan(&n)

	ketemu := false
	*i = searchIdx(*A, *jumlahData, n)

	if *i > -1 {
		fmt.Println("\033[32mData ditemukan..\033[0m")
		headingTable()
		cetak(*A, *i)
		footerTable()
		ketemu = true
	}
	if !ketemu {
		fmt.Println("\033[31mMaaf, data tidak ditemukan\033[0m")
	}
	fmt.Println()
}

// Fungsi untuk mencari indeks data transaksi berdasarkan ID transaksi
// Fungsi ini menggunakan binary search untuk mencari ID transaksi
func searchIdx(A dataTransaksi, jD int, n string) int {
	var low, mid, high, idx int

	low = 0
	high = jD - 1
	idx = -1

	for low <= high && idx == -1 {
		mid = (low + high) / 2
		if A[mid].idTransaksi == n {
			idx = mid
		} else if A[mid].idTransaksi < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

// Fungsi untuk mencari data transaksi berdasarkan nama barang
// Nama barang yang dicari dimasukkan oleh pengguna
// Fungsi ini menggunakan sequential search untuk mencari nama barang
func cariNamaBarang(A *dataTransaksi, jumlahData *int) {
	var n string
	var ketemu bool = false
	fmt.Print("Masukan Nama Barang (contoh: Tempat_Pensil): ")
	fmt.Scan(&n)

	for i := 0; i < *jumlahData; i++ {
		if n == A[i].namaBarang {
			if !ketemu {
				fmt.Println("\033[32mData ditemukan..\033[0m")
				headingTable()
			}
			cetak(*A, i)
			ketemu = true
		}
	}

	if ketemu {
		footerTable()
	} else {
		fmt.Println("\033[31mMaaf, data tidak ditemukan\033[0m")
	}
	jeda()
	fmt.Println()
}

// Fungsi untuk mengedit data transaksi
func editData(A *dataTransaksi, jumlahData, idx *int) {
	var namaBarang string
	var jumlah, totalHarga int

	cariIdTransaksi(A, jumlahData, idx)
	if *idx > -1 {
		fmt.Println("Silakan edit data")
		fmt.Scan(&namaBarang, &jumlah, &totalHarga)
		if konfirmasi() == true {
			A[*idx].namaBarang = namaBarang
			A[*idx].jumlah = jumlah
			A[*idx].totalHarga = totalHarga
			fmt.Printf("\033[32mData ke %d sudah berhasil di edit\n\033[0m", *idx+1)
		}
	}
	jeda()
	fmt.Println()
}

// Fungsi untuk menghapus data transaksi
func hapus(A *dataTransaksi, jumlahData, idx *int) {
	cariIdTransaksi(A, jumlahData, idx)
	if *idx > -1 && konfirmasi() == true {
		for i := *idx; i < *jumlahData-1; i++ {
			A[i] = A[i+1]
		}
		A[*jumlahData-1] = data{}

		*jumlahData--
		fmt.Printf("\033[32mData ke %d sudah terhapus\n\033[0m", *idx+1)
	}
	jeda()
	fmt.Println()
}

// Fungsi untuk menghapus semua data transaksi (reset)
func reset(A *dataTransaksi, jumlahData *int) {
	fmt.Println("Data yang akan dihapus: ")
	cetakData(*A, *jumlahData)

	if konfirmasi() == true {
		for i := 0; i < *jumlahData; i++ {
			A[i].idTransaksi = "0"
			A[i].namaBarang = "0"
			A[i].jumlah = 0
			A[i].totalHarga = 0
		}
		*jumlahData = 0
		fmt.Println("\033[32mSemua data sudah terhapus\033[0m")
	}
	jeda()
	fmt.Println()
}

// Fungsi untuk menampilkan menu konfirmasi
func konfirmasi() bool {
	var pilihan string
	var k bool = false
	fmt.Print("\033[33mApakah anda yakin akan mengedit atau menghapus data ini? (y/n)\033[0m")
	fmt.Scan(&pilihan)
	fmt.Println()

	if pilihan == "y" || pilihan == "Y" {
		k = true
	}
	return k
}

//Menu kopweasi simpan pinjam
func menuSimpanPinjam(){
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║      K O P E R A S I  S I M P A N  P I N J A M     ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║      1. Input Data Transaksi                       ║")
	fmt.Println("║      2. Cetak Data Transaksi                       ║")
	fmt.Println("║      3. Search                                     ║")
	fmt.Println("║      4. Edit Data                                  ║")
	fmt.Println("║      5. Hapus Data                                 ║")
	fmt.Println("║      6. Reset Data                                 ║")
	fmt.Println("║      7. Exit                                       ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Print("Pilih (1/2/3/4/5/6/7)? ")
}


func simpanPinjam(B *simpanPinjamTransaksi, jumlahDataSP *int) {
	var pilihan, idx int

	for pilihan < 7 {
		clearScreen()
		menuSimpanPinjam()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			bacaSP(B, jumlahDataSP)
		case 2:
			cetakMenuSP(*B, *jumlahDataSP)
		case 3:
			searchSP(B, jumlahDataSP, &idx)
		case 4:
			editDataSP(B, jumlahDataSP, &idx)
		case 5:
			hapusSP(B, jumlahDataSP, &idx)
		case 6:
			resetSP(B, jumlahDataSP)
		}
	}
}

// Fungsi untuk membaca data transaksi dari pengguna
// Data yang dimasukkan adalah nama pelaku, status, dan total transaksi
func bacaSP(B *simpanPinjamTransaksi, jumlahData *int) {
	var namaPelaku, status string
	var totalTransaksi int
	var stop bool = false
	fmt.Println("Silakan masukan 0 0 0 jika ingin berhenti")
	fmt.Println("Nama barang | Status | Total Harga")
	for !stop {
		fmt.Scan(&namaPelaku, &status, &totalTransaksi)
		if namaPelaku == "0" && status == "0" && totalTransaksi == 0 {
			stop = true
		} else if totalTransaksi > 0 && (status == "simpan" || status == "pinjam") {
			B[*jumlahData].idTransaksi = fmt.Sprintf("KSP%02d", *jumlahData+1) // penginputan id transaksi dilakukan secara otomatis
			B[*jumlahData].namaPelaku = namaPelaku
			B[*jumlahData].status = status
			B[*jumlahData].totalTransaksi = totalTransaksi
			*jumlahData++
		} else {
			fmt.Println("\033[31mTotal transaksi harus > 0 dan status harus 'simpan' atau 'pinjam'\033[0m")
		}
	}
	fmt.Println("\033[32mData sudah berhasil diinput\033[0m")
	jeda()
	fmt.Println()
}

// Fungsi untuk menampilkan menu cetak
func cetakMenuSP(B simpanPinjamTransaksi, jumlahData int) {
	var pilihan int
	for pilihan < 4 {
		clearScreen()
		menuCetakSP()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			cetakDataSP(B, jumlahData)
		case 2:
			terurutMenurunSP(B, jumlahData)
		case 3:
			terurutMenaikSP(B, jumlahData)
		}
	}
}

func menuCetakSP() {
    fmt.Println("╔════════════════════════════════════════════════════╗")
    fmt.Println("║               CETAK DATA SIMPAN PINJAM             ║")
    fmt.Println("╠════════════════════════════════════════════════════╣")
    fmt.Println("║      1. Tampilkan Semua Data Transaksi             ║")
    fmt.Println("║      2. Tampilkan Data Transaksi Terurut Menurun   ║")
    fmt.Println("║      3. Tampilkan Data Transaksi Terurut Menaik    ║")
    fmt.Println("║      4. Kembali ke Menu Utama                      ║")
    fmt.Println("╚════════════════════════════════════════════════════╝")
    fmt.Print("Pilih (1/2/3/4)? ")
}

// fungsi untuk mencetak header dan footer tabel
func headingTableSP() {
    fmt.Println("╔══════════════╦══════════════╦════════════╦══════════════════╗")
    fmt.Printf("║ %-12s ║ %-12s ║ %-10s ║ %-16s ║\n", "ID Transaksi", "Nama Pelaku", "Status", "Total Transaksi")
    fmt.Println("╠══════════════╬══════════════╬════════════╬══════════════════╣")
}

func footerTableSP() {
	fmt.Println("╚══════════════════════════════════════════╩══════════════════╝")
}

// fungsi untuk mencetak data transaksi
func cetakSP(B simpanPinjamTransaksi, i int) {
    fmt.Printf("║ %-12s ║ %-12s ║ %-10s ║ %-16d ║\n", 
        B[i].idTransaksi, B[i].namaPelaku, B[i].status, B[i].totalTransaksi)
}

// fungsi untuk mencetak semua data transaksi dan menjumlahkan nilai total simpan dan pinjam 
func cetakDataSP(B simpanPinjamTransaksi, jumlahData int) {
    clearScreen()
    if B[0] == (dataSimpanPinjam{}) {
        fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
    } else {
        headingTableSP()
        
        totalSimpan := 0
        totalPinjam := 0
        for i := 0; i < jumlahData; i++ {
            cetakSP(B, i)
            if B[i].status == "simpan" {
                totalSimpan += B[i].totalTransaksi
            } else {
                totalPinjam += B[i].totalTransaksi
            }
        }
        cetakTotalNilaiTransaksiSP(totalSimpan, totalPinjam)
        footerTableSP()
    }
    jeda()
    fmt.Println()
}

// fungsi untuk mencetak total nilai transaksi
// total nilai transaksi dihitung dengan menjumlahkan total simpan dan total pinjam
func cetakTotalNilaiTransaksiSP(totalSimpan, totalPinjam int) {
    fmt.Println("╠══════════════╩══════════════╩════════════╬══════════════════╣")
    fmt.Printf("║ %-40s ║ %-16d ║\n", "Total Simpan", totalSimpan)
    fmt.Println("╠══════════════════════════════════════════╬══════════════════╣")
    fmt.Printf("║ %-40s ║ %-16d ║\n", "Total Pinjam", totalPinjam)
}

func menuKriteriaSortSP() int {
    var kriteria int
    fmt.Println("Urut berdasarkan:")
    fmt.Println("1. Nama Pelaku")
    fmt.Println("2. Total Transaksi")
    fmt.Print("Pilih (1/2): ")
    fmt.Scan(&kriteria)
    return kriteria
}

// Fungsi untuk mengurutkan data transaksi secara menurun
// Fungsi ini menggunakan insertion sort untuk mengurutkan data berdasarkan kriteria yang dipilih
func terurutMenurunSP(B simpanPinjamTransaksi, jumlahData int) {
    clearScreen()
    if jumlahData == 0 {
        fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
        jeda()
    } else {
        var C simpanPinjamTransaksi
        C = B
        kriteria := menuKriteriaSortSP()
        insertionSortDescSP(&C, jumlahData, kriteria)

        fmt.Println("\033[32mData terurut menurun:\033[0m")
        cetakDataSP(C, jumlahData)
    }
}

func insertionSortDescSP(B *simpanPinjamTransaksi, jumlahData int, kriteria int) {
    for i := 1; i < jumlahData; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && ((kriteria == 1 && B[j].namaPelaku < key.namaPelaku) || 
                       (kriteria == 2 && B[j].totalTransaksi < key.totalTransaksi)) {
            B[j+1] = B[j]
            j--
        }
        B[j+1] = key
    }
}

// Fungsi untuk mengurutkan data transaksi secara menaik
// Fungsi ini menggunakan selection sort untuk mengurutkan data berdasarkan kriteria yang dipilih
func terurutMenaikSP(B simpanPinjamTransaksi, jumlahData int) {
    clearScreen()
    if jumlahData == 0 {
        fmt.Println("\033[31mTidak ada data untuk ditampilkan\033[0m")
        jeda()
    } else {
        var C simpanPinjamTransaksi
        C = B
        kriteria := menuKriteriaSortSP()
        selectionSortSP(&C, jumlahData, kriteria)

        fmt.Println("\033[32mData terurut menaik:\033[0m")
        cetakDataSP(C, jumlahData)
    }
}

func selectionSortSP(B *simpanPinjamTransaksi, jumlahData int, kriteria int) {
    for i := 0; i < jumlahData-1; i++ {
        minIdx := i
        for j := i + 1; j < jumlahData; j++ {
            if (kriteria == 1 && B[j].namaPelaku < B[minIdx].namaPelaku) ||
               (kriteria == 2 && B[j].totalTransaksi < B[minIdx].totalTransaksi) {
                minIdx = j
            }
        }
        B[i], B[minIdx] = B[minIdx], B[i]
    }
}

// Fungsi untuk mencari data transaksi berdasarkan ID transaksi atau nama pelaku
func searchSP(B *simpanPinjamTransaksi, jumlahData, i *int) {
    var pilihan int

    for pilihan < 3 {
        clearScreen()
        menuSearchSP()
        fmt.Scan(&pilihan)
        fmt.Println()
        switch pilihan {
        case 1:
            cariIdTransaksiSP(B, jumlahData, i)
            jeda()
        case 2:
            cariNamaPelakuSP(B, jumlahData)
        }
    }
}

// Fungsi untuk menampilkan menu pencarian
func menuSearchSP() {
    fmt.Println("╔════════════════════════════════════════════════════╗")
    fmt.Println("║               PENCARIAN SIMPAN PINJAM              ║")
    fmt.Println("╠════════════════════════════════════════════════════╣")
    fmt.Println("║      1. Berdasarkan ID Transaksi                   ║")
    fmt.Println("║      2. Berdasarkan Nama Pelaku                    ║")
    fmt.Println("║      3. Exit                                       ║")
    fmt.Println("╚════════════════════════════════════════════════════╝")
    fmt.Print("Pilih (1/2/3)? ")
}

// Fungsi untuk mencari data transaksi berdasarkan ID transaksi
// ID transaksi yang dicari dimasukkan oleh pengguna
func cariIdTransaksiSP(B *simpanPinjamTransaksi, jumlahData, i *int) {
    var n string

    fmt.Print("Masukan ID Transaksi (contoh: KSP01): ")
    fmt.Scan(&n)

    ketemu := false
    *i = searchIdxSP(*B, *jumlahData, n)

    if *i > -1 {
        fmt.Println("\033[32mData ditemukan..\033[0m")
        headingTableSP()
        cetakSP(*B, *i)
        footerTableSP()
        ketemu = true
    }
    if !ketemu {
        fmt.Println("\033[31mMaaf, data tidak ditemukan\033[0m")
    }
    fmt.Println()
}

// Fungsi untuk mencari indeks data transaksi berdasarkan ID transaksi
// Fungsi ini menggunakan binary search untuk mencari ID transaksi
func searchIdxSP(B simpanPinjamTransaksi, jD int, n string) int {
    var low, mid, high, idx int

    low = 0
    high = jD - 1
    idx = -1

    for low <= high && idx == -1 {
        mid = (low + high) / 2
        if B[mid].idTransaksi == n {
            idx = mid
        } else if B[mid].idTransaksi < n {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return idx
}

// Fungsi untuk mencari data transaksi berdasarkan nama Pelaku
// Nama Pelaku yang dicari dimasukkan oleh pengguna
// Fungsi ini menggunakan sequential search untuk mencari nama barang
func cariNamaPelakuSP(B *simpanPinjamTransaksi, jumlahData *int) {
    var n string
    var ketemu bool = false
    fmt.Print("Masukan Nama Pelaku: ")
    fmt.Scan(&n)

    for i := 0; i < *jumlahData; i++ {
        if n == B[i].namaPelaku {
            if !ketemu {
                fmt.Println("\033[32mData ditemukan..\033[0m")
                headingTableSP()
            }
            cetakSP(*B, i)
            ketemu = true
        }
    }

    if ketemu {
        footerTableSP()
    } else {
        fmt.Println("\033[31mMaaf, data tidak ditemukan\033[0m")
    }
    jeda()
    fmt.Println()
}

// Fungsi untuk mengedit data transaksi
func editDataSP(B *simpanPinjamTransaksi, jumlahData, idx *int) {
    var namaPelaku, status string
    var totalTransaksi int

    cariIdTransaksiSP(B, jumlahData, idx)
    if *idx > -1 {
        fmt.Println("Silakan edit data")
        fmt.Scan(&namaPelaku, &status, &totalTransaksi)
        if konfirmasi() == true {
            B[*idx].namaPelaku = namaPelaku
            B[*idx].status = status
            B[*idx].totalTransaksi = totalTransaksi
            fmt.Printf("\033[32mData ke %d sudah berhasil di edit\n\033[0m", *idx+1)
        }
    }
    jeda()
    fmt.Println()
}

// Fungsi untuk menghapus data transaksi
func hapusSP(B *simpanPinjamTransaksi, jumlahData, idx *int) {
    cariIdTransaksiSP(B, jumlahData, idx)
    if *idx > -1 && konfirmasi() == true {
        for i := *idx; i < *jumlahData-1; i++ {
            B[i] = B[i+1]
        }
        B[*jumlahData-1] = dataSimpanPinjam{}

        *jumlahData--
        fmt.Printf("\033[32mData ke %d sudah terhapus\n\033[0m", *idx+1)
    }
    jeda()
    fmt.Println()
}

// Fungsi untuk menghapus semua data transaksi (reset)
func resetSP(B *simpanPinjamTransaksi, jumlahData *int) {
    fmt.Println("Data yang akan dihapus: ")
    cetakDataSP(*B, *jumlahData)

    if konfirmasi() == true {
        for i := 0; i < *jumlahData; i++ {
            B[i].idTransaksi = "0"
            B[i].namaPelaku = "0"
            B[i].status = "0"
            B[i].totalTransaksi = 0
        }
        *jumlahData = 0
        fmt.Println("\033[32mSemua data sudah terhapus\033[0m")
    }
    jeda()
    fmt.Println()
}

// penutup
func hehe(){
	 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡏⡀⠀⢄⠈⢙⢒⠲⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⠶⠚⠃⢣⠀⠸⣶⠋⢈⠇⣿⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⢀⣠⠤⢄⡀⠀⠀⢀⣤⣶⢖⣿⣏⡷⠰⠞⠉⠉⢺⣇⠀⢹⠀⢸⡀⠋⢿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⣠⠤⠶⣚⣋⡅⠀⠀⠙⣻⡟⠋⡔⠛⠉⠁⠙⣷⠀⠰⣿⡄⠀⢻⡆⠀⢧⠀⠙⢦⠀⡿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⢸⣯⡵⠛⠛⠉⠈⣷⠀⠘⠉⠀⠙⢦⡀⠀⣾⣆⠀⢻⣇⠀⢳⣳⡀⠘⣷⣀⣨⠙⢲⣒⣤⢣⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⢸⡿⡇⢴⠀⠀⣿⠈⣇⠀⢸⣣⡀⠘⣧⡀⠙⠿⠀⠈⣿⣄⣸⡏⢻⣾⣥⢨⣿⣿⠿⠛⠛⠉⠀⠀⣀⣤⣤⠤⠤⣒⡲⠤⠤⣤⣤⣤⣤⣀⡀⢧⡀⢹⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠙⠻⡆⢧⠀⠹⡄⢹⡄⠘⢏⢧⠀⢸⡷⣤⣤⣶⣚⠋⢸⠛⢳⣿⡅⠘⢇⢷⣿⠀⠀⣤⢔⣻⡯⠶⠛⠛⠋⠉⠉⠉⠉⣉⣉⣛⡛⠓⢲⣬⡝⣷⣤⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⢿⡘⡆⠀⢻⠀⢳⣄⣼⠛⢳⣾⢥⣴⠛⠉⠁⠉⢻⡄⠈⢣⢣⠀⠸⡎⠈⣧⣯⡾⠋⠁⠀⣠⣤⣴⣾⣿⣿⣿⣿⠿⢛⣭⣴⣿⡿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⢧⡛⠀⠚⢰⣿⣿⠛⠳⣼⡇⠀⣿⠀⠘⣟⡄⠈⢳⡀⠘⠿⠃⢠⡇⣠⣿⠛⢀⣤⣾⣿⣿⣿⣿⣿⡿⠟⣫⣤⣞⡯⠚⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠉⠛⠪⠖⢻⣝⣦⠀⠘⠇⠀⡟⣦⠀⠛⠿⠀⡼⠓⡦⠤⠴⣫⣤⣿⡁⣠⣿⣿⣿⡿⠟⣡⣴⣿⣳⣛⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠤⠒⢦⣀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢎⠳⡄⠀⠀⡇⢈⣓⣦⡴⢛⡥⣺⡿⠿⢿⣿⣿⠁⣵⣿⣿⣿⡿⠟⣡⣴⣿⣳⣛⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠤⠒⢦⣀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢳⠈⠆⠀⡧⢸⣿⠛⠛⠛⠛⠁⠀⣠⣾⣿⠃⣼⣿⣿⣿⢋⣴⠞⠋⠀⠴⠿⠿⣿⣯⣹⣀⣀⣀⡀⠀⠀⠀⠀⠀⣘⣦⣤⣀⠈⠑⢦⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣇⠀⣴⣃⡬⣿⣁⡀⠀⠀⠀⣼⣿⢿⡇⡴⠋⠀⣻⡿⠛⠁⠀⠀⠀⠀⠀⠀⠉⢙⣶⡿⠿⠿⠭⣍⣒⣶⢤⣀⣏⣀⠉⠓⣝⡄⠈⢧⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠲⢷⡿⠋⢉⣿⠀⢀⣞⡿⠃⠸⠃⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⢦⣤⣀⣀⡀⠈⠉⠳⢮⣝⠯⣶⡀⠈⣿⠦⠼⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⣰⡟⠀⣰⡿⠋⢹⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⣿⣿⣿⣿⣿⣷⣶⣄⡀⠉⠳⣝⢿⡉⠁⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣾⣽⣿⣿⣏⣿⠀⠀⣿⠁⢠⣿⢻⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣝⠿⣿⣿⣿⣿⣿⣿⣷⣄⠘⢷⣙⣄⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⣠⣾⡿⠞⠉⠸⢏⣾⡟⠉⢷⡀⣿⣦⣼⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿⣶⣌⡙⠿⣿⣿⣿⣿⣿⣷⣄⠻⣍⢦⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⢸⣿⡇⠀⠈⢉⣿⢿⣿⣫⣿⠀⠀⠀⠀⠀⣀⠀⠀⠀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣱⠁⠙⠯⢿⡶⣤⣉⡻⠿⣿⣿⣿⣦⠙⣮⢦⠀⠀⠀")
    fmt.Println("⠀⠀⠀⣴⣿⠟⠀⠀⠀⠀⠀⠀⠘⣟⣇⣀⣶⡯⠟⠋⠉⠀⢹⣆⠀⠀⠀⠀⠻⣦⣄⣀⣤⡶⠀⠀⠀⠀⠀⢐⠶⢀⣴⣿⠁⠀⠀⠀⠀⠉⠙⠯⢽⣷⡶⣍⣙⠻⣷⠈⢯⣣⡀⠀")
    fmt.Println("⠀⠀⡼⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠹⣾⣿⣋⠀⠀⠀⠀⠀⠀⢻⣆⠀⠀⠀⠀⠈⠉⠉⠁⠀⠀⠀⢀⣀⠀⠘⠷⠛⣿⣹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠙⠻⢿⣶⣤⣜⢯⣣⡀")
    fmt.Println("⠀⢸⣿⠃⢰⠀⣀⠀⠀⠀⠀⠀⠀⠀⠘⣿⠖⠀⠀⠀⠀⠀⠀⠀⢹⣷⣤⡀⠀⠀⠀⠀⣷⣄⣀⣀⣈⡉⣀⡀⠀⢠⡿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠿⣟⠇⠇")
    fmt.Println("⢀⣿⡟⠀⠀⠀⣿⠀⢰⣇⠀⡴⠂⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⣰⠏⠀⠉⠛⠶⣤⣤⡀⠀⠉⠉⠁⠈⠙⢋⣠⣴⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠁")
    fmt.Println("⢸⣿⡇⠀⠀⠀⠻⣄⣸⠙⢦⡀⠀⠀⢸⣇⠀⠀⠀⠀⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⠉⠛⠛⠛⠛⠛⠛⢿⣷⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⢸⣿⡇⠀⠀⠀⠀⠙⠋⠀⠀⠹⠀⠀⢸⡿⣦⠀⠀⠀⠀⠀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⡇⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⠁⠈⢷⣄⢀⣄⢰⣷⢠⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡟⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⢹⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡿⠀⠀⠀⠙⠟⣿⣿⠻⣿⣇⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⢷⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠘⣧⠀⠀⠀⠀⠀⠲⣿⠀⠉⠹⣿⣦⠀⠀⠀⠀⠀⠀⢀⡀⢀⣰⣷⠎⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠙⢟⣷⣄⡀⠀⠀⠀⠀⠀⠀⢻⡆⠀⠀⠀⠀⠀⢹⣇⠀⠀⠈⢙⣿⣷⡀⠀⢀⣴⣿⣴⣟⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠙⠿⢟⣶⣤⣄⣀⣀⣀⣨⣿⣦⣀⣀⣀⣀⣀⣿⡄⠀⠀⠈⠉⣿⣿⡶⠛⠉⠀⣾⡝⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⠒⠚⠛⠓⠊⠉⠿⣿⣿⡋⢹⡁⣸⢷⡄⠀⠀⠀⠸⣆⣀⠀⠀⣾⣽⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠑⠯⠭⠭⠭⢽⣿⣦⡀⠰⣄⢺⣿⣆⣾⣽⣻⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
    fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠚⠛⠲⣻⣛⣿⡿⣿⣾⡏⠀")
}