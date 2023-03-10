package main

type user struct {
	id       int
	username int
	password int
}

// Username dan password seharusnya bertipe data string, bukan integer
// Tipe data pada password yang dikelompokkan ke integer tidak aman dalam penyimpanan

type userservice struct {
	t []user
}

// Terdapat variabel t disitu yang sifatnya ambigu/tidak deskriptif

func (u userservice) getallusers() []user {
	return u.t
}

// Di dalam method userservice variabel struct-nya tidak digunakan, namun mengambil data yang telah disimpan di luar struct-nya

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	// Pada method getusertbyid disitu, jika id tidak menemukan user yang sesuai nantinya kode mengembalikan user yang kosong, sebaiknya ditambahkan mekanisme error handling

	return user{}
}
