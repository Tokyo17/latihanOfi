package task

import "fmt"

func Kasir() {
	listharga := map[string]int{
		"Milo":    5000,
		"Es Krim": 1000,
		"Mie":     2000,
	}
	lisbarang := []string{}
	var barang string
	var jumlah int
	var totalHarga int
	var selesai string
	var uangCustomer int

	for {
		fmt.Print("masukan barang : ")
		fmt.Scan(&barang)
		lisbarang = append(lisbarang, barang)
		fmt.Print("masukan Jumlah barang : ")
		fmt.Scan(&jumlah)
		totalHarga += jumlah * listharga[barang]
		fmt.Print("Sudah selesai ? [Y/N]: ")
		fmt.Scan(&selesai)
		if selesai == "y" {
			break
		}
	}
	fmt.Println(lisbarang)
	fmt.Println("Total Belnja anda : ", totalHarga)
	fmt.Print("masukan uang anda : ")
	fmt.Scan(&uangCustomer)
	fmt.Println("Kembalian anda : ", uangCustomer-totalHarga)
}
