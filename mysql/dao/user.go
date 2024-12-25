package dao

import (
	"easy-storage/mysql/models"
	"gorm.io/gorm"
	"log"
)

// GetUserList 获取用户的基础数据列表
func GetUserList(db *gorm.DB) (err error, users []*models.User) {
	err = db.Table((&models.User{}).Table()).Find(&users).Error
	if err != nil {
		log.Fatalf("find user list error [%v]", err)
	}
	return err, users
}
