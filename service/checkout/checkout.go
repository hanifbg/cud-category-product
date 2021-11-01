package checkout

import "time"

type Checkout struct {
	ID             uint
	CartID         uint
	UserID         uint
	Subtotal       int
	Payment        string
	IsActive       bool
	StatusPayment  bool
	StatusDelivery bool
	StatusOverall  bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewCheckout(cartID int,
	userID int,
	subTotal int,
	payment string,
	isActive bool,
	statusPayment bool,
	statusDelivery bool,
	statusOverall bool,
	createdAt time.Time,
	updatedAt time.Time,
) Checkout {
	return Checkout{
		ID:             0,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		CartID:         uint(cartID),
		UserID:         uint(userID),
		Subtotal:       subTotal,
		Payment:        payment,
		IsActive:       isActive,
		StatusPayment:  statusPayment,
		StatusDelivery: statusDelivery,
		StatusOverall:  statusOverall,
	}
}

func (oldData *Checkout) ModifyCheckout(
	newActive bool,
	newStatusPay bool,
	newStatusDeliv bool,
	newStatusOver bool,
) Checkout {
	return Checkout{
		ID:             oldData.ID,
		CreatedAt:      oldData.CreatedAt,
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
		CartID:         oldData.CartID,
		UserID:         oldData.UserID,
		Subtotal:       oldData.Subtotal,
		Payment:        oldData.Payment,
		IsActive:       newActive,
		StatusPayment:  newStatusPay,
		StatusDelivery: newStatusDeliv,
		StatusOverall:  newStatusOver,
	}
}
