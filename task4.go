package main

import "fmt"

type Sepeda struct {
	jumlahBan  int
	jumlahGigi int
	warna      string
}

func (s *Sepeda) waktuTempuh(jarak float32) float32 {
	waktu := jarak * 2.5
	return waktu
}

func main() {
	sepeda := make(map[int]Sepeda)
	var jarak float32

	sepeda[1] = Sepeda{2, 6, "Merah"}
	sepeda[2] = Sepeda{2, 4, "Hijau"}
	sepeda[3] = Sepeda{2, 4, "Biru"}
	sepeda[4] = Sepeda{2, 4, "Pink"}
	sepeda[5] = Sepeda{2, 4, "Coklat"}

	for index, element := range sepeda {
		fmt.Println(element)
		fmt.Print("Masukan Jarak Tempuh : ")
		fmt.Scanln(&jarak)
		sepedaKe := sepeda[index]
		fmt.Println("Waktu Tempuh Dengan Jarak ", jarak, " Km Adalah ", sepedaKe.waktuTempuh(jarak), " menit")
		fmt.Println("---------------------------")

	}
}
