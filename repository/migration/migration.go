package migration

import (
	"github.com/hanifbg/cud-category-product/repository/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
