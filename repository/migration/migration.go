package migration

import (
	"github.com/hanifbg/cud-category-product/repository/cart"
	"github.com/hanifbg/cud-category-product/repository/category"
	"github.com/hanifbg/cud-category-product/repository/product"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&category.Category{}, &product.Product{}, &cart.Cart{}, &cart.CartProduct{})
}
