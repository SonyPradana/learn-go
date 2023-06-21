package main

import (
	"fmt"
)

func main() {
	// var
	var name string = "prada"
	age := 12
	var keren = true

	// condition
	if keren {
		fmt.Printf("its work! %s %d \n", name, age)

		switch {
		case (age < 17):
			fmt.Println("masih muda ya")
		}
	}

	// loop
	fmt.Println("pass:")
	for i := 0; i < age; i++ {
		fmt.Printf("saat usia saya %d saya...\n", i)
	}

	// keren juga go
	// var short int8 = 0

	// for short < 123 {
	// 	fmt.Println(short)
	// 	short++
	// }

	// array
	var service = [...]string{
		"layanan",
		"informasi",
		"BPJS",
		"PBI",
	}

	fmt.Println(service[0], service[1], service[2], service[3])

	var wow = [2][2]string{
		{"aku", "keren"},
		{"aku2", "keren2"},
	}

	for _, value := range wow {
		for _, isi := range value {
			fmt.Println(isi)
		}
	}

	// Slice - reference array
	var langs = [...]string{"vb.net", "c#", "php", "js", "rust"}

	var alang = langs[0:3]
	var blang = langs[2:3]
	var aalang = alang[0:3]
	var bblang = blang[2:3]

	fmt.Println(langs)
	fmt.Println(alang)
	fmt.Println(blang)
	fmt.Println(aalang)
	fmt.Println(bblang)

	// willbe effect with first referens
	bblang[0] = "go"

	fmt.Println(langs)
	fmt.Println(alang)
	fmt.Println(blang)
	fmt.Println(aalang)
	fmt.Println(bblang)

	// menghitung lebar
	var asli = [...]string{"vb.net", "c#", "php", "js", "rust"}
	var sebelum = asli[0:4]

	fmt.Println("sebelum cap -", cap(sebelum))
	fmt.Println("sebelum len - ", len(sebelum))

	sebelum = append(sebelum, "go")
	fmt.Println(asli)
	fmt.Println(sebelum)

	fmt.Println("sebelum cap -", cap(sebelum))
	fmt.Println("sebelum len - ", len(sebelum))

	sebelum = append(sebelum, "c")

	fmt.Println("sebelum cap -", cap(sebelum))
	fmt.Println("sebelum len - ", len(sebelum))

	fmt.Println(asli)
	fmt.Println(sebelum)

	// belum paham
	// var makanan = []string{"soto", "gorengan"}
	// var minuman = []string{"es teh", "air putih"}
	// var order = copy(makanan, minuman)

	// fmt.Println(makanan)
	// fmt.Println(minuman)
	// fmt.Println(order)
}
