package product

import (
	"time"

	"github.com/hanifbg/cud-category-product/service/product"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint       `gorm:"id;primaryKey;autoIncrement"`
	CategoryID uint       `gorm:"category_id"`
	Name       string     `gorm:"name"`
	Price      int        `gorm:"price"`
	Stock      int        `gorm:"stock"`
	Image      string     `gorm:"image"`
	Detail     string     `gorm:"detail"`
	CreatedAt  time.Time  `gorm:"created_at"`
	UpdatedAt  time.Time  `gorm:"updated_at"`
	DeletedAt  *time.Time `gorm:"deleted_at"`
}

func (col *Product) ToProduct() product.Product {
	var product product.Product

	product.ID = col.ID
	product.CategoryID = col.CategoryID
	product.Name = col.Name
	product.Price = col.Price
	product.Stock = col.Stock
	product.Image = col.Image
	product.Detail = col.Detail
	product.CreatedAt = col.CreatedAt
	product.UpdatedAt = col.UpdatedAt
	product.DeletedAt = col.DeletedAt

	return product
}

func newProductTable(product product.Product) *Product {
	return &Product{
		product.ID,
		product.CategoryID,
		product.Name,
		product.Price,
		product.Stock,
		product.Image,
		product.Detail,
		product.CreatedAt,
		product.UpdatedAt,
		product.DeletedAt,
	}
}

type GormRepository struct {
	DB *gorm.DB
}

func (repo *GormRepository) FindProductById(id int) (*product.Product, error) {

	var productData Product

	err := repo.DB.First(&productData, id).Error
	if err != nil {
		return nil, err
	}

	product := productData.ToProduct()

	return &product, nil
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) AddProduct(product product.Product) error {
	productData := newProductTable(product)

	err := repo.DB.Create(productData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormRepository) UpdateProduct(product product.Product) error {
	categoryData := newProductTable(product)

	err := repo.DB.Model(&categoryData).Where("category_id = ?", product.CategoryID).Updates(Product{
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Price:      product.Price,
		Stock:      product.Stock,
		Image:      product.Image,
		Detail:     product.Detail,
		UpdatedAt:  product.UpdatedAt,
		DeletedAt:  product.DeletedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
