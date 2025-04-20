package models

type representasiData struct {
    Users Users
    Products Products
    Restaurant Restaurant
    Warehouse Warehouse
}

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


type Products struct {
    ID          string  // ID unik produk
    Name        string  // Nama menu (misal: Nasi Goreng)
    Description string  // Deskripsi singkat
    Category    string  // Misal: makanan, minuman, snack, dessert
    Price       float64 // Harga satuan
    Stock       int     // Jumlah stok tersedia
    ImageURL    string  // (optional) URL ke gambar menu
}



type Restaurant struct {
    ID          string  // ID unik restoran
    Name        string  // Nama restoran
    Address     string  // Alamat lengkap restoran
    PhoneNumber string  // Nomor telepon restoran
    Email       string  // Email restoran (opsional)
    OpeningHours string // Jam operasional (misal: 08:00 - 22:00)
    Rating      float64 // Rating rata-rata restoran (opsional)
}


type Warehouse struct {
    ID          string   // ID unik gudang
    Name        string   // Nama gudang (misal: Gudang Toko, Gudang Pasar)
    Location    string   // Lokasi gudang
    Capacity    int      // Kapasitas maksimum gudang (jumlah item)
    CurrentStock int     // Jumlah stok saat ini di gudang
    ManagerName string   // Nama pengelola gudang
    Contact     string   // Kontak pengelola gudang
    Type        string   // Tipe gudang (misal: "Toko" atau "Pasar")
    Storage     []string // Daftar item yang disimpan (misal: ikan, daging, sabun, dll.)
}