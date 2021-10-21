package migration

import (
	"github.com/hanifbg/login_register_v2/repository/user"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
