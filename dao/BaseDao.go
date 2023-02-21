package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

// InitDb 初始化数据库连接 main.go中调用
func InitDb() {
	//dataSource := fmt.Sprintf("%s:%s@(%s)/%s", conf.Config.Mysql.User, conf.Config.Mysql.Password, conf.Config.Mysql.Host, conf.Config.Mysql.DbName)
	dataSource := "root:141097@(localhost:3306)/jmall"
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		fmt.Printf("sqlx open failed,error:%v\n", err)
		return
	}
	Db = db
}

// CloseDb main.go中调用
func CloseDb() {
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
