package main

import "fmt"

func main() {
	siswa := map[string]string{
		"nama_siswa": "Ade Rifqy Aulian",
		"kelas":      "XI RPL 2",
		"hobi":       "Coding",
	}

	fmt.Println(siswa)
	fmt.Println(siswa["nama_siswa"])
	fmt.Println(siswa["kelas"])

	motor := make(map[string]string)
	motor["merk"] = "Yamaha"
	motor["warna"] = "Hitam"

	motor["tahun"] = "2020"

	fmt.Println(motor)

	buah := map[string]string{
		"a": "apel",
		"b": "pisang",
		"c": "jeruk",
	}

	fmt.Println(buah)
	fmt.Println(len(buah))

	game := map[string]string{
		"judul": "Mobile Legends",
		"genre": "MOBA",
		"mode":  "Ranked",
	}

	fmt.Println(game)
	delete(game, "mode")
	fmt.Println(game)
	fmt.Println(len(game))

	akun := map[string]string{
		"username": "ade_rifqy",
		"role":     "user",
	}

	akun["role"] = "admin"
	fmt.Println(akun)

	profil := map[string]string{
		"nama_profil": "Ade Rifqy Aulian",
		"alamat":      "Talang Pantai",
		"pekerjaan":   "Programmer",
	}

	fmt.Println(profil["nama_profil"])
	profil["umur"] = "20"
	delete(profil, "alamat")
	fmt.Println(profil)
	fmt.Println(len(profil))

}
