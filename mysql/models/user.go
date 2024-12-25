package models

//type User struct {
//	ID           int64  `gorm:"primaryKey"`
//	UserName     string `gorm:"column:user_name" json:"user_name"`
//	UserAccount  string `gorm:"column:user_account" json:"user_account"`
//	UserPassword string `gorm:"column:user_password" json:"user_password"`
//	Status       int8   `gorm:"column:status" json:"status"`
//	CreatedTime  int64  `gorm:"column:created_time" json:"created_time"`
//	UpdatedTime  int64  `gorm:"column:updated_time" json:"updated_time"`
//	IsDeleted    int8   `gorm:"column:is_deleted" json:"is_deleted"`
//}

// User  用户的表，这里是个模型
type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName     string `gorm:"size:128;not null;default:'';uniqueIndex" json:"user_name"`    // 用户姓名
	UserAccount  string `gorm:"size:128;not null;default:'';uniqueIndex" json:"user_account"` // 用户账号
	UserPassword string `gorm:"size:1024;not null;default:''" json:"user_password"`           // 用户密码
	Status       int    `gorm:"type:tinyint;not null;default:0" json:"status"`                // 是否活跃
	CreatedTime  int    `gorm:"not null" json:"created_time"`                                 // 创建时间
	UpdatedTime  int    `gorm:"default:0" json:"updated_time"`                                // 更新时间
	IsDeleted    int    `gorm:"type:tinyint;not null;default:0" json:"is_deleted"`            // 是否已经删除
}

func (u *User) Table() string {
	return "user"
}
