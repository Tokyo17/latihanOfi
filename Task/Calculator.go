package task

import "fmt"

func Calculator() {
	var angka1, angka2 int
	var perintah string
	fmt.Print("masukan angka 1 : ")
	fmt.Scan(&angka1)
	fmt.Print("masukan angka 2 : ")
	fmt.Scan(&angka2)
	fmt.Print("masukan perintah : ")
	fmt.Scan(&perintah)

	fmt.Print("hasilnya : ")
	if perintah == "tambah" {
		fmt.Println(angka1 + angka2)
	} else if perintah == "kurang" {
		fmt.Println(angka1 - angka2)
	} else if perintah == "bagi" {
		fmt.Println(angka1 / angka2)
	} else if perintah == "kali" {
		fmt.Println(angka1 * angka2)
	}
}
