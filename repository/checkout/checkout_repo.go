package checkout

import (
	"time"

	"github.com/hanifbg/cud-category-product/service/checkout"
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

type Checkout struct {
	ID             uint       `gorm:"id;primaryKey;autoIncrement"`
	CartID         uint       `gorm:"cart_id"`
	UserID         uint       `gorm:"user_id"`
	Subtotal       int        `gorm:"subtotal"`
	Payment        string     `gorm:"payment"`
	IsActive       bool       `gorm:"is_active"`
	StatusPayment  bool       `gorm:"status_payment"`
	StatusDelivery bool       `gorm:"status_delivery"`
	StatusOverall  bool       `gorm:"status_overall"`
	CreatedAt      time.Time  `gorm:"created_at"`
	UpdatedAt      time.Time  `gorm:"updated_at"`
	DeletedAt      *time.Time `gorm:"deleted_at"`
}

func newCheckoutTable(checkout checkout.Checkout) *Checkout {
	return &Checkout{
		checkout.ID,
		checkout.CartID,
		checkout.UserID,
		checkout.Subtotal,
		checkout.Payment,
		checkout.IsActive,
		checkout.StatusPayment,
		checkout.StatusDelivery,
		checkout.StatusOverall,
		checkout.CreatedAt,
		checkout.UpdatedAt,
		checkout.DeletedAt,
	}
}

func (col *Checkout) ToCheckout() checkout.Checkout {
	var checkout checkout.Checkout

	checkout.ID = col.ID
	checkout.CreatedAt = col.CreatedAt
	checkout.UpdatedAt = col.UpdatedAt
	checkout.DeletedAt = col.DeletedAt
	checkout.UserID = col.UserID
	checkout.CartID = col.CartID
	checkout.Subtotal = col.Subtotal
	checkout.Payment = col.Payment
	checkout.IsActive = col.IsActive
	checkout.StatusPayment = col.StatusPayment
	checkout.StatusDelivery = col.StatusDelivery
	checkout.StatusOverall = col.StatusOverall

	return checkout
}

func (repo *GormRepository) CreateCheckout(checkout checkout.Checkout) error {
	newCheckout := newCheckoutTable(checkout)
	result := repo.DB.Create(newCheckout)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) FindCheckoutById(id int) (*checkout.Checkout, error) {

	var checkoutData Checkout

	err := repo.DB.First(&checkoutData, id).Error
	if err != nil {
		return nil, err
	}

	checkout := checkoutData.ToCheckout()

	return &checkout, nil
}

func (repo *GormRepository) UpdateCheckout(checkout checkout.Checkout) error {
	checkoutData := newCheckoutTable(checkout)

	err := repo.DB.Model(&checkoutData).Where("id = ?", checkout.ID).Updates(Checkout{
		IsActive:       checkout.IsActive,
		StatusPayment:  checkout.StatusPayment,
		StatusDelivery: checkout.StatusDelivery,
		StatusOverall:  checkout.StatusOverall,
		DeletedAt:      checkout.DeletedAt,
	}).Error

	if err != nil {
		return err
	}
	return nil
}
