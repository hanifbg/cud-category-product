package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string `json:"name"  validate:"required"`
	Email        string `json:"email" validate:"required,email" gorm:"type:varchar(20)"`
	Phone_number string `json:"phone_number" validate:"required,number"`
	Password     string `json:"password"  validate:"required"`
	Address      string `json:"address"  validate:"required"`
	Role         int
	Token_hash   string
}

type GormRepository struct {
	DB *gorm.DB
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//InsertUser Insert new User into storage
func (repo *GormRepository) RepoFuncForUser() error {

	return nil
}
