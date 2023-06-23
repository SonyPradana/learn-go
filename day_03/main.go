package main

import (
	"fmt"
)

func main() {
	var names = []string{"taylor", "otwell"}
	print_nama(names...)

	var short_func = func() {
		fmt.Println("short funct")
	}

	short_func()

	func(time int) {
		fmt.Println("dua ditulis", time)
	}(2)

	// kemablian sebagai fungsi
	var kembali = kembalian(1)
	var nilai int = kembali(2)
	fmt.Println(nilai) // 3
}

func print_nama(names ...string) {
	for _, name := range names {
		fmt.Println("nama: ", name)
	}
}

func kembalian(awal int) func(int) int {
	return func(bisakah int) int {
		return 0 + bisakah + awal
	}
}
