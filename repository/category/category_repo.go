package category

import (
	"time"

	"github.com/hanifbg/cud-category-product/service/category"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint       `gorm:"id;primaryKey;autoIncrement"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	Name      string     `json:"name"  validate:"required"`
	IsActive  bool       `gorm:"default:true"`
}

func (col *Category) ToCategory() category.Category {
	var category category.Category

	category.ID = col.ID
	category.CreatedAt = col.CreatedAt
	category.UpdatedAt = col.UpdatedAt
	category.DeletedAt = col.DeletedAt
	category.Name = col.Name
	category.IsActive = col.IsActive

	return category
}

func newCategoryTable(category category.Category) *Category {
	return &Category{
		category.ID,
		category.CreatedAt,
		category.UpdatedAt,
		category.DeletedAt,
		category.Name,
		category.IsActive,
	}
}

type GormRepository struct {
	DB *gorm.DB
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) AddCategory(category category.Category) error {
	categoryData := newCategoryTable(category)

	err := repo.DB.Create(categoryData).Error
	if err != nil {
		return err
	}
	return nil
}
