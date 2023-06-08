package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	// Substitusi cipher
	for i := 0; i < len(s.name); i++ {
		if s.name[i] >= 'a' && s.name[i] <= 'z' {
			nameEncode += string('a' + (s.name[i]-'a'+3)%26) // geser 3 karakter
		} else if s.name[i] >= 'A' && s.name[i] <= 'Z' {
			nameEncode += string('A' + (s.name[i]-'A'+3)%26) // geser 3 karakter
		} else {
			nameEncode += string(s.name[i]) // karakter lainnya tidak diubah
		}
	}

	s.nameEncode = nameEncode // simpan hasil enkripsi pada properti nameEncode
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	// Substitusi cipher
	for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i] >= 'a' && s.nameEncode[i] <= 'z' {
			nameDecode += string('a' + (s.nameEncode[i]-'a'+23)%26) // geser ke kiri 3 karakter
		} else if s.nameEncode[i] >= 'A' && s.nameEncode[i] <= 'Z' {
			nameDecode += string('A' + (s.nameEncode[i]-'A'+23)%26) // geser ke kiri 3 karakter
		} else {
			nameDecode += string(s.nameEncode[i]) // karakter lainnya tidak diubah
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode) // input yang digunakan untuk dekripsi adalah hasil enkripsi sebelumnya
		fmt.Print("\nDecode of student’s name " + a.nameEncode + "is : " + c.Decode())
	}
}
