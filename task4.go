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

	sepeda[1] = Sepeda{2, 6, "Merah"}
	sepedaPertama := sepeda[1]
	fmt.Println(sepedaPertama)
	fmt.Println("Waktu Tempuh Dengan Jarak 20km Adalah", sepedaPertama.waktuTempuh(20.0))
	fmt.Println("-----------------------------")

	sepeda[2] = Sepeda{2, 4, "Hijau"}
	sepedaKedua := sepeda[2]
	fmt.Println(sepedaKedua)
	fmt.Println("Waktu Tempuh Dengan Jarak 40km Adalah", sepedaKedua.waktuTempuh(40.0))
	fmt.Println("-----------------------------")

	sepeda[3] = Sepeda{2, 4, "Biru"}
	sepedaKetiga := sepeda[3]
	fmt.Println(sepedaKetiga)
	fmt.Println("Waktu Tempuh Dengan Jarak 100km Adalah", sepedaKetiga.waktuTempuh(100.0))
	fmt.Println("-----------------------------")

	sepeda[4] = Sepeda{2, 4, "Pink"}
	sepedaKeempat := sepeda[4]
	fmt.Println(sepedaKeempat)
	fmt.Println("Waktu Tempuh Dengan Jarak 200km Adalah", sepedaKeempat.waktuTempuh(200.0))
	fmt.Println("-----------------------------")

	sepeda[5] = Sepeda{2, 4, "Coklat"}
	sepedaKelima := sepeda[5]
	fmt.Println(sepedaKelima)
	fmt.Println("Waktu Tempuh Dengan Jarak 1000km Adalah", sepedaKelima.waktuTempuh(1000.0))

}
