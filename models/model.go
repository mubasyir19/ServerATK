package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID
	Username    string
	Password    string
	FullName    string
	Address     string
	PhoneNumber string
	Orders      []Order `gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model
	ID       uuid.UUID
	Name     string
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Stock       int64
	Photo       string
	CategoryID  uuid.UUID
}

type Order struct {
	gorm.Model
	ID           uuid.UUID
	UserID       uuid.UUID
	Date         time.Time
	TotalPrice   float64
	AlamatKirim  string
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	gorm.Model
	ID        uuid.UUID
	OrderID   uuid.UUID
	ProductID uuid.UUID
	Quantity  int64
	UnitPrice int64
	SubTotal  int64
}

type Cart struct {
	gorm.Model
	ID     uuid.UUID
	UserID uuid.UUID
}

type StatusBayar string

const (
	StatusPending StatusBayar = "Pending"
	StatusLunas   StatusBayar = "Success"
	StatusGagal   StatusBayar = "Failed"
)

type Payment struct {
	gorm.Model
	ID       uuid.UUID
	TotalPay float64
	PayDate  time.Time
	Status   StatusBayar
}
