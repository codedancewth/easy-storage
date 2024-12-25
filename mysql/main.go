package main

import (
	"easy-storage/mysql/dao"
	"easy-storage/mysql/tool"
	"fmt"
	"gorm.io/gorm"
	"log"
)

// 此处进行mysql的主函数的方式
func main() {

	// 初始化数据库
	mysql := (&Mysql{UserName: UserName, PassWord: PassWord}).InitMysql()

	// case 1 查询用户的数据
	// DescribeUserInfo(mysql)

	// case 2 迁移用户的数据，不存在则生成，只会新增，不回减少，亲测有效
	AutoMigrate(mysql)
}

func DescribeUserInfo(db *gorm.DB) {
	// 查询用户的列表
	err, users := dao.GetUserList(db)
	if err != nil {
		log.Panicf("err [%v]", err)
		return
	}
	// 打印查询结果
	fmt.Println("Query results:")
	for _, user := range users {
		fmt.Printf("user [%+v]", user)
	}
}

func AutoMigrate(db *gorm.DB) {
	tool.AutoMigrate(db)
}
