package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID       uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama     string `gorm:"type:varchar(255);column:nama"`
	Kategori string `gorm:"type:varchar(50);column:kategori"`
	Harga    int    `gorm:"type:int;column:harga"`
}

func (Product) TableName() string {
	return "produck2"
}

type Penduduk struct {
	ID     uint     `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Alamat []Alamat `gorm:"many2many:penduduk_alamat"`
}

type Alamat struct {
	ID            uint   `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	AlamatLengkap string `gorm:"column:alamat_lengkap"`
}

func main() {
	connURI := "postgresql://citizix_user:S3cret@localhost:5432/citizix_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Printf("Failed connect db: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("Connected to db")
	db.AutoMigrate(&Product{}, &Alamat{}, &Penduduk{})

	// produk := Product{Nama: "Kertas A4", Kategori: "Pakaian", Harga: 10000}
	// result := db.Create(&produk)
	// if result.Error != nil {
	// 	fmt.Printf("Failed fill table: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println("Data added")

	// var produk Product
	// result := db.First(&produk, 2)
	// produk := Product{ID: 1}
	// result := db.First(&produk)
	// if result.Error != nil {
	// 	fmt.Printf("Failed get table: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println(produk)

	// var productSlice []Product
	// result := db.Find(&productSlice)
	// if result.Error != nil {
	// 	fmt.Printf("Failed get table: %v\n", result.Error)
	// 	os.Exit(1)
	// }

	// fmt.Println(productSlice)

	// var productSlice []Product
	// result := db.Find(&productSlice, []uint{1, 2})
	// if result.Error != nil {
	// 	fmt.Printf("Failed get table: %v\n", result.Error)
	// 	os.Exit(1)
	// }

	// fmt.Println(productSlice)

	// var productSlice []Product
	// result := db.Where(map[string]interface{}{"id": 3}).Find(&productSlice)
	// if result.Error != nil {
	// 	fmt.Printf("Failed get table: %v\n", result.Error)
	// 	os.Exit(1)
	// }

	// fmt.Println(productSlice)

	// var productSlice []Product
	// result := db.Not(map[string]interface{}{"id": 3}).Find(&productSlice)
	// if result.Error != nil {
	// 	fmt.Printf("Failed get table: %v\n", result.Error)
	// 	os.Exit(1)
	// }

	// fmt.Println(productSlice)

	// result := db.Model(&Product{ID: 1}).Updates(&Product{Nama: "Update", Harga: 90000})
	// if result.Error != nil {
	// 	fmt.Printf("Failed update data: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println("Data updated")

	// result := db.Delete(Product{ID: 1})
	// if result.Error != nil {
	// 	fmt.Printf("Failed delete data: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println("Data deleted")

	// db.Transaction(func(tx *gorm.DB) error {
	// 	if result := tx.Delete(&Product{ID: 1}); result.Error != nil {
	// 		fmt.Printf("Failed transaction: %v\n", result.Error)
	// 		return result.Error
	// 	} else {
	// 		fmt.Println("Transaction successfully")
	// 		return nil
	// 	}
	// })

	// penduduk := Penduduk{
	// 	Alamat: []Alamat{
	// 		{
	// 			AlamatLengkap: "Jakarta",
	// 		},
	// 		{
	// 			AlamatLengkap: "Bandung",
	// 		},
	// 	},
	// }

	// db.Create(&penduduk)

	var penduduk Penduduk
	db.Preload("Alamat").First(&penduduk, 1)
	fmt.Println(penduduk)
}
