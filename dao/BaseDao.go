package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"mall/conf"
	"mall/model"
	"time"
)

var Db *gorm.DB

// InitDb 初始化数据库连接 main.go中调用
func InitDb() {
	dataSource := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", conf.MysqlConf.User, conf.MysqlConf.Password, conf.MysqlConf.Host, conf.MysqlConf.DbName)

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		fmt.Printf("gorm open failed,error:%v\n", err)
		return
	}
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	Db = db
	//生成的数据库表明不带s
	db.SingularTable(true)
	//自动迁移
	db.AutoMigrate(&model.User{}, &model.Address{}, &model.GoodsImage{},
		&model.Goods{}, &model.Category1{}, &model.Category2{}, &model.Category3{},
		&model.GoodsBanner{}, &model.PromotionBanner{}, &model.OrderGoods{},
		&model.OrderInfo{}, &model.Role{}, &model.Menu{}, &model.RoleMenu{},
		&model.Member{})

}

// CloseDb main.go中调用
func CloseDb() {
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Paginate 分页
func Paginate(current, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if current == 0 {
			current = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (current - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
