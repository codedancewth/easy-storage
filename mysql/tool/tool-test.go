package tool

import (
	"gorm.io/gorm"
	"log"
)

// 发现有些db的函数挺有趣的，收藏一下

// AutoMigrate 自动迁移模式（如果表不存在，会自动创建）
func AutoMigrate(db *gorm.DB) {
	//type UserV2 struct {
	//	ID           int64  `gorm:"primaryKey"`
	//	UserName     string `gorm:"column:user_name" json:"user_name"`
	//	UserAccount  string `gorm:"column:user_account" json:"user_account"`
	//	UserPassword string `gorm:"column:user_password" json:"user_password"`
	//	Status       int8   `gorm:"column:status" json:"status"`
	//	CreatedTime  int64  `gorm:"column:created_time" json:"created_time"`
	//	UpdatedTime  int64  `gorm:"column:updated_time" json:"updated_time"`
	//	IsDeleted    int8   `gorm:"column:is_deleted" json:"is_deleted"`
	//}

	type UserV2 struct {
		ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
		UserName     string `gorm:"size:128;not null;default:'';uniqueIndex" json:"user_name"`    // 用户姓名
		UserAccount  string `gorm:"size:128;not null;default:'';uniqueIndex" json:"user_account"` // 用户账号
		UserPassword string `gorm:"size:1024;not null;default:''" json:"user_password"`           // 用户密码
		Status       int    `gorm:"type:tinyint;not null;default:0" json:"status"`                // 是否活跃
		CreatedTime  int    `gorm:"not null" json:"created_time"`                                 // 创建时间
		UpdatedTime  int    `gorm:"default:0" json:"updated_time"`                                // 更新时间
		IsDeleted    int    `gorm:"type:tinyint;not null;default:0" json:"is_deleted"`            // 是否已经删除
	}

	// 迁移，存在不管，否则创建新的表
	err := db.AutoMigrate(&UserV2{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
		return
	}

	log.Println("AutoMigrate successful!")
}
