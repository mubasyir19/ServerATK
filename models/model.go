package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Username    string
	Password    string
	NamaLengkap string
	Alamat      string
	NoTelpon    string
	Pesanan     Pesanan
	Keranjang   Keranjang
}

type Kategori struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	Nama   string
	Produk []Produk // One-to-many => 1 kategori punya banyak produk
}

type Produk struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid"`
	Nama           string
	Deskripsi      string
	Harga          string
	Stok           int64
	Foto           string
	KategoriProduk uint
}

type Pesanan struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	UserID     uint
	Tanggal    time.Time `gorm:"type:datetime"`
	HargaTotal int64
	// StatusPesanan
	AlamatKirim string
	ItemPesanan []ItemPesanan
	Pembayaran  Pembayaran
}

type ItemPesanan struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	OrderID     uint
	ProdukID    uint
	Jumlah      int64
	HargaSatuan int64
	SubTotal    int64
}

type Keranjang struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	UserID uint
}

type StatusBayar string

const (
	StatusPending StatusBayar = "Pending"
	StatusLunas   StatusBayar = "Lunas"
	StatusGagal   StatusBayar = "Gagal"
)

type Pembayaran struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	// MetodeBayar
	TotalBayar   int64
	TanggalBayar time.Time `gorm:"type:datetime"`
	Status       StatusBayar
}
