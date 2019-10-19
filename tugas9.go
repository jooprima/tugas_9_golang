package main

import "fmt"
import "strconv"

func main() {
	defer pulihkan_saya()
	var input string
	fmt.Print("masukan angka :")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "ini adalah angka")
	} else {
		fmt.Println(input, "ini bukan angka")
		panic(err.Error())
		fmt.Println("munculkan saya")
	}
}

func pulihkan_saya() {
	if r := recover(); r != nil {
		fmt.Println("ini adalah huruf atau kata!!!")
	} else {
		fmt.Println("tepat sekali!!!")
	}
}
