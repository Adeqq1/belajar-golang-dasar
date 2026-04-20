package main

import "fmt"

func main() {
	name := "Ade"

	if name == "Ade" {
		fmt.Println("Halo Ade")
	} else if name == "Budi" {
		fmt.Println("Halo Budi")
	} else {
		fmt.Println("Halo Stranger")
	}

	umur := 17
	if umur >= 18 {
		fmt.Println("Sudah Dewasa")
	} else {
		fmt.Println("Masih Remaja")
	}

	nilai := 85
	if nilai >= 90 {
		fmt.Println("Grade A")
	} else if nilai >= 80 {
		fmt.Println("Grade B")
	} else if nilai >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	password := "admin123"

	if password == "admin123" {
		fmt.Println("Login Berhasil")
	} else {
		fmt.Println("Password Salah")
	}

	username := "aderifqy"

	if length := len(username); length > 10 {
		fmt.Println("Username terlalu panjang")
	} else {
		fmt.Println("Username valid")
	}

	angka := 3

	if angka > 0 {
		fmt.Println("Bilangan Positif")
	} else if angka == 0 {
		fmt.Println("Bilangan nol")
	} else {
		fmt.Println("Bilangan negatif")
	}

	nama := "Ade"
	umur = 20

	if nama == "Ade" && umur >= 18 {
		fmt.Println("Selamat Datang Ade Dewasa")
	} else if nama == "Ade" && umur <= 18 {
		fmt.Println("Selamata Daang Ade Remaja")
	} else {
		fmt.Println("User tidak dikenal")
	}

	nilai = 78

	if nilai >= 90 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	username = "ade"

	if username == "ade" {
		fmt.Println("Selamat datang admin")
	} else {
		fmt.Println("Akses user biasa")
	}

	angka = 0

	if angka > 0 {
		fmt.Println("Positif")
	} else if angka == 0 {
		fmt.Println("Nol")
	} else {
		fmt.Println("Negatif")
	}

}
