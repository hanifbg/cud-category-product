package cart

import "time"

type Cart struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	UserID     uint
	TotalPrice int
}

func NewCart(userId int, totalPrice int, createdAt time.Time, updatedAt time.Time) Cart {
	return Cart{
		ID:         0,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		UserID:     uint(userId),
		TotalPrice: totalPrice,
	}
}

func (oldCart *Cart) ModifyCart(newTotal int, updated time.Time) Cart {
	return Cart{
		ID:         oldCart.ID,
		UserID:     oldCart.UserID,
		TotalPrice: newTotal,
		CreatedAt:  oldCart.CreatedAt,
		UpdatedAt:  updated,
	}
}

type CartProduct struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	CartID    uint
	ProductID uint
	Count     int
	Price     int
}

func NewCartProduct(
	cartId int,
	productID int,
	count int,
	price int,
	createdAt time.Time,
	updatedAt time.Time,
) CartProduct {
	return CartProduct{
		ID:        0,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CartID:    uint(cartId),
		ProductID: uint(productID),
		Count:     count,
		Price:     price,
	}
}

func (oldCartProduct *CartProduct) ModifyCartProduct(newPrice int, newCount int, updated time.Time) CartProduct {
	return CartProduct{
		ID:        oldCartProduct.ID,
		CartID:    oldCartProduct.CartID,
		ProductID: oldCartProduct.ProductID,
		CreatedAt: oldCartProduct.CreatedAt,
		UpdatedAt: updated,
		Count:     newCount,
		Price:     newPrice,
	}
}
