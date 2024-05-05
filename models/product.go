package models

import (
	"log/slog"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Title       string `json:"title" gorm:"column:title"`
	Price       string `json:"price" gorm:"column:price"`
	Size        string `json:"size" gorm:"column:size"`
	Measurement string `json:"measurement" gorm:"mesurement"`
	Brand       string `json:"brand" gorm:"column:brand"`
	Quantity    string `json:"quantity" gorm:"column:quantity"`
}

func GetProductByID(ID uint, db *gorm.DB) (*Product, error) {
	var product Product
	if result := db.Where("ID=?", ID).First(&product); result.Error != nil {
		slog.Info("Error fetching product", "Error", result.Error.Error())
		return nil, result.Error
	}
	return &product, nil
}

func DeleteProduct(ID uint, db *gorm.DB) error {
	product, err := GetProductByID(ID, db)
	if err != nil {
		return err
	}

	if result := db.Delete(&product); result.Error != nil {
		slog.Info("Error deleting product", "Error", result.Error.Error())
		return err
	}
	return nil
}

func GetAllProduct(db *gorm.DB) (*[]Product, error) {
	var products []Product
	if result := db.Find(&products); result.Error != nil {
		slog.Info("Error fetching products", "Error", result.Error.Error())
		return nil, result.Error
	}
	return &products, nil
}
