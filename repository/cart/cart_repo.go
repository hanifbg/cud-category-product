package cart

import (
	"time"

	"github.com/hanifbg/cud-category-product/service/cart"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

type Cart struct {
	ID         uint       `gorm:"id;primaryKey;autoIncrement"`
	UserID     uint       `gorm:"user_id"`
	TotalPrice int        `gorm:"total_price"`
	CreatedAt  time.Time  `gorm:"created_at"`
	UpdatedAt  time.Time  `gorm:"updated_at"`
	DeletedAt  *time.Time `gorm:"deleted_at"`
}

func newCartTable(cart cart.Cart) *Cart {
	return &Cart{
		cart.ID,
		cart.UserID,
		cart.TotalPrice,
		cart.CreatedAt,
		cart.UpdatedAt,
		cart.DeletedAt,
	}
}

func (col *Cart) ToCart() cart.Cart {
	var cart cart.Cart

	cart.ID = col.ID
	cart.CreatedAt = col.CreatedAt
	cart.UpdatedAt = col.UpdatedAt
	cart.DeletedAt = col.DeletedAt
	cart.UserID = col.UserID
	cart.TotalPrice = col.TotalPrice

	return cart
}

func (repo *GormRepository) FindCartByUserID(id int) (*cart.Cart, error) {
	var cartData Cart

	err := repo.DB.Where("user_id = ?", id).First(&cartData).Error
	if err != nil {
		return nil, err
	}

	cart := cartData.ToCart()

	return &cart, nil
}

func (repo *GormRepository) CreateCart(cart cart.Cart) (*cart.Cart, error) {
	newCart := newCartTable(cart)
	result := repo.DB.Create(newCart)
	err := result.Error
	if err != nil {
		return nil, err
	}

	cartData := newCart.ToCart()

	return &cartData, nil
}

func (repo *GormRepository) UpdateCart(cart cart.Cart) (*cart.Cart, error) {

	cartData := newCartTable(cart)
	result := repo.DB.Model(&cartData).Updates(map[string]interface{}{
		"total_price": cart.TotalPrice,
		"updated_at":  cart.UpdatedAt,
	})
	err := result.Error
	if err != nil {
		return nil, err
	}

	updatedCart := cartData.ToCart()

	return &updatedCart, nil
}
