package models

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

type Users struct {
    ID       string `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    BisLoc   string `json:"bis_loc"`
    DateLoc  string `json:"date_loc"`
    Year     string `json:"year"`
    Role     string `json:"role"`
}

type Products struct {
    ID          string  `gorm:"primaryKey" json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Category    string  `json:"category"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    ImageURL    string  `json:"image_url"`
}

type OrderItem struct {
    Produ   Products `json:"products"`
    Value int `json:"value"`
    Subtotal  int    `json:"subtotal"`
}

type Pesnan struct {
    ID          uint         `gorm:"primaryKey" json:"id"`
    Products    []OrderItem  `gorm:"-" json:"products"`     // tidak masuk DB
    ProductsRaw string       `json:"products_raw"`          // masuk DB
    Table       string       `json:"table_id"`
    Waiter      string       `json:"waiter_id"`
    Time        string       `json:"time"`
}
