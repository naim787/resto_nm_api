package models

type Users struct {
	Name string // id users
	Id string // id users
	Email string // email users
	Password string // password users
	Bis_Loc string // lokasi bisnis
	Date_Loc string // data users di buat
	Year string // taggal lahir
	Role string // manager, users
}


type Product struct {
    ID          string  // ID unik produk
    Name        string  // Nama menu (misal: Nasi Goreng)
    Description string  // Deskripsi singkat
    Category    string  // Misal: makanan, minuman, snack, dessert
    Price       float64 // Harga satuan
    Stock       int     // Jumlah stok tersedia
    ImageURL    string  // (optional) URL ke gambar menu
}