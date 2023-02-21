package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"mall/model"
	"time"
)

var Db *gorm.DB

// InitDb 初始化数据库连接 main.go中调用
func InitDb() {
	//todo 从App.yml文件中读取配置信息
	//dataSource := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true", conf.Config.User, conf.Config.Password, conf.Config.Host, conf.Config.DbName)
	dataSource := "root:141097@(localhost:3306)/mall?charset=utf8mb4&parseTime=true"
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
	db.SingularTable(true)
	//自动迁移
	db.AutoMigrate(&model.User{}, &model.Address{}, &model.GoodsImage{},
		&model.GoodsSku{}, &model.GoodsSpu{}, &model.GoodsType{},
		&model.GoodsBanner{}, &model.TypeGoodsBanner{},
		&model.PromotionBanner{}, &model.OrderGoods{}, &model.OrderInfo{})
}

// CloseDb main.go中调用
func CloseDb() {
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
