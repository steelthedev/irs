package models

import (
	"log/slog"

	"github.com/steelthedev/irs/data"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Title       string  `json:"title" gorm:"column:title"`
	Price       float64 `json:"price" gorm:"column:price"`
	Size        string  `json:"size" gorm:"column:size"`
	Measurement string  `json:"measurement" gorm:"mesurement"`
	Brand       string  `json:"brand" gorm:"column:brand"`
	Quantity    int     `json:"quantity" gorm:"column:quantity"`
}

type ProductService struct {
	DB *gorm.DB
}

func (p *Product) GetTotalQuantityPrice() int {
	return (int(p.Price) * p.Quantity)
}

func (ps *ProductService) AddNewProduct(productParam data.AddProductParams) (*Product, error) {

	product := Product{
		Title:       productParam.Title,
		Price:       productParam.Price,
		Size:        productParam.Size,
		Measurement: productParam.Measurement,
		Brand:       productParam.Brand,
		Quantity:    productParam.Quantity,
	}

	if result := ps.DB.Create(&product); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (ps *ProductService) GetProductByID(ID uint) (*Product, error) {
	var product Product
	if result := ps.DB.Where("ID=?", ID).First(&product); result.Error != nil {
		slog.Info("Error fetching product", "Error", result.Error.Error())
		return nil, result.Error
	}
	return &product, nil
}

func (ps *ProductService) DeleteProduct(ID uint) error {
	product, err := ps.GetProductByID(ID)
	if err != nil {
		return err
	}

	if result := ps.DB.Delete(&product); result.Error != nil {
		slog.Info("Error deleting product", "Error", result.Error.Error())
		return err
	}
	return nil
}

func (ps *ProductService) GetAllProduct() (*[]Product, error) {
	var products []Product
	if result := ps.DB.Find(&products); result.Error != nil {
		slog.Info("Error fetching products", "Error", result.Error.Error())
		return nil, result.Error
	}
	return &products, nil
}

func (ps *ProductService) RemoveQuantity(ID uint, units int) error {

	product, err := ps.GetProductByID(ID)
	if err != nil {
		return err
	}
	product.Quantity -= units
	if result := ps.DB.Save(&product); result.Error != nil {
		slog.Info("Error updating product quantity", "Error", result.Error.Error())
		return err
	}
	return nil
}

func (ps *ProductService) AddQuantity(ID uint, units int) error {

	product, err := ps.GetProductByID(ID)
	if err != nil {
		return err
	}
	product.Quantity += units
	if result := ps.DB.Save(&product); result.Error != nil {
		slog.Info("Error updating product quantity", "Error", result.Error.Error())
		return err
	}
	return nil
}
