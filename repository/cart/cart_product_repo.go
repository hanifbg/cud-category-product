package cart

import (
	"time"

	"github.com/hanifbg/cud-category-product/service/cart"
)

type CartProduct struct {
	ID        uint       `gorm:"id;primaryKey;autoIncrement"`
	CartID    uint       `gorm:"cart_id"`
	ProductID uint       `gorm:"product_id"`
	Count     int        `gorm:"count"`
	Price     int        `gorm:"price"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}

type SumPrice struct {
	SumPrice int `gorm:"sum_price"`
}

func newCartProductTable(cart cart.CartProduct) *CartProduct {
	return &CartProduct{
		cart.ID,
		cart.CartID,
		cart.ProductID,
		cart.Count,
		cart.Price,
		cart.CreatedAt,
		cart.UpdatedAt,
		cart.DeletedAt,
	}
}

func (col *CartProduct) ToCartProduct() cart.CartProduct {
	var cart cart.CartProduct

	cart.ID = col.ID
	cart.CreatedAt = col.CreatedAt
	cart.UpdatedAt = col.UpdatedAt
	cart.DeletedAt = col.DeletedAt
	cart.CartID = col.CartID
	cart.ProductID = col.ProductID
	cart.Count = col.Count
	cart.Price = col.Price

	return cart
}

func (repo *GormRepository) FindCartProduct(cartID int, productID int) (*cart.CartProduct, error) {
	var cartProductData CartProduct

	err := repo.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartProductData).Error
	if err != nil {
		return nil, err
	}

	cartProduct := cartProductData.ToCartProduct()

	return &cartProduct, nil
}

func (repo *GormRepository) CreateCartProduct(cartProduct cart.CartProduct) (*cart.CartProduct, error) {
	newCartProduct := newCartProductTable(cartProduct)
	result := repo.DB.Create(newCartProduct)
	err := result.Error
	if err != nil {
		return nil, err
	}
	cartProductData := newCartProduct.ToCartProduct()
	return &cartProductData, nil
}

func (repo *GormRepository) UpdateCartProduct(cartProduct cart.CartProduct) (*cart.CartProduct, error) {

	cartProductData := newCartProductTable(cartProduct)
	result := repo.DB.Model(&cartProductData).Updates(map[string]interface{}{
		"count":      cartProduct.Count,
		"price":      cartProduct.Price,
		"updated_at": cartProduct.UpdatedAt,
	})
	err := result.Error
	if err != nil {
		return nil, err
	}

	updatedCartProduct := cartProductData.ToCartProduct()

	return &updatedCartProduct, nil
}

func (repo *GormRepository) SumPrice(cartID int) int {
	var result SumPrice
	repo.DB.Table("cart_products").Select("sum(price) as sum_price").Where("cart_id = ?", cartID).Scan(&result)

	return result.SumPrice
}
