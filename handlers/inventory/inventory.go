package inventory

import "gorm.io/gorm"

type InventoryHandler struct {
	DB *gorm.DB
}
