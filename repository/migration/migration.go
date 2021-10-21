package migration

import (
	"GO_MOD_MODULE_NAME/repository/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
