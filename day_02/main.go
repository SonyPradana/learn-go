package main

import (
	"fmt"
)

func main() {
	// map
	var user = map[string]string{
		"name": "taylor",
		"age":  "fourty two",
		"lang": "php",
	}

	for key, val := range user {
		fmt.Println(key, " -> ", val)
	}

	var profile = make(map[string]string)

	profile["user"] = "taylor"
	profile["real_name"] = "taylor otwell"
	profile["des"] = "creator and mainter of laravel framework"

	fmt.Println("user: ", profile["user"])
	fmt.Println("real_name: ", profile["real_name"])

	// profile["region"] = "UE"
	var region, have_ragion = profile["region"]

	if have_ragion {
		fmt.Println(region)
	}

	// map dan slice
	var users = []map[string]string{
		{"user": "taylor", "product": "laravel", "prefix": "mr."},
		{"user": "sony", "product": "savanna", "prefix": "mr."},
	}

	for _, user := range users {
		fmt.Println("")
		fmt.Println("user", user["user"])
		fmt.Println("product", user["product"])
	}

	print_product("taylor", users)

	fmt.Println(plus(1, 2))

	var sebuah_string, sebuah_int = string_dan_int()

	fmt.Println(sebuah_string, sebuah_int)

	var _, umur = predefine_return()

	fmt.Println(umur)
}

func print_product(user string, users []map[string]string) {
	fmt.Println("searching for:", user)
	for _, item := range users {
		if item["user"] == user {
			fmt.Println(getPrintedProduct(item))
		}
	}
}

func getPrintedProduct(users map[string]string) string {
	return "find " + users["product"] + ", by " + users["prefix"] + " " + users["user"]
}

// parameter sama ditulis di ahir
func plus(a, b int) int {
	return a + b
}

// go dapat me-return dua hasil
func string_dan_int() (string, int) {
	return "abc", 123
}

func predefine_return() (nama string, umur int) {
	nama = "taylor"
	umur = 43

	return
}
