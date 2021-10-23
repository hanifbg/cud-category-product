package product

import "time"

type Product struct {
	ID         uint
	CategoryID uint
	Name       string
	Price      int
	Stock      int
	Image      string
	Detail     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewProduct(categoryID int,
	name string,
	price int,
	stock int,
	image string,
	detail string,
	createdAt time.Time,
	updatedAt time.Time,
) Product {
	return Product{
		ID:         0,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		CategoryID: uint(categoryID),
		Name:       name,
		Price:      price,
		Stock:      stock,
		Image:      image,
		Detail:     detail,
	}
}

func (oldData *Product) ModifyProduct(
	newCategoryId int,
	newName string,
	newPrice int,
	newStock int,
	newImage string,
	newDetail string,
	deletedAt *time.Time,
) Product {
	return Product{
		ID:         oldData.ID,
		CategoryID: uint(newCategoryId),
		Name:       newName,
		Price:      newPrice,
		Stock:      newStock,
		Image:      newImage,
		Detail:     newDetail,
		UpdatedAt:  time.Now(),
		DeletedAt:  deletedAt,
	}
}
