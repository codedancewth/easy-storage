package main

import (
	"easy-storage/mysql/dao"
	"easy-storage/mysql/models"
	"easy-storage/mysql/tool"
	"fmt"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

// 此处进行mysql的主函数的方式
func main() {

	// 初始化数据库
	mysql := (&Mysql{UserName: UserName, PassWord: PassWord}).InitMysql()

	// case 1 查询用户的数据
	//DescribeUserInfo(mysql)

	// case 2 迁移用户的数据，不存在则生成，只会新增，不回减少，亲测有效
	//AutoMigrate(mysql)

	// case 3 自动生成月表
	AutoBuildMonthTable(mysql)

	// case 4 模拟并发连接数
	//AutoDescribeUser(mysql)

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

func AutoBuildMonthTable(db *gorm.DB) {
	// 判断当前是否是这个月的第一天
	if time.Now().Day() != 1 {
		fmt.Println(fmt.Sprintf("this day is %d", time.Now().Day()))
		return
	}

	// TODO 可以设置缓存来判断了第一天的情况

	// 生成当前月的表名
	tableName := tool.GetMonthlyTableName("user")

	fmt.Printf("table name %s", tableName)
	// 确保表存在
	err := tool.EnsureMonthlyTable(db, &models.User{}, tableName)

	if err != nil {
		log.Fatalf("failed to ensure table: %v", err)
	}
}

func AutoDescribeUser(db *gorm.DB) {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DescribeUserInfo(db)
		}()
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
