package main

import "fmt"

var pl = fmt.Println

func main() {
	pl("Belajar variabel")

	var name string = "Dava aditya perdana"
	var umur uint8 = 19
	var address string = "Medan"
	pl("Nama saya adalah :", name)
	pl("Umur saya adalah :", umur)
	pl("Alamat saya adalah :", address)

}
