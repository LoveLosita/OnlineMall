package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/OnlineMall?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
	//第一部分：连接数据库，并检测其连接正常性
	var err error
	Db, err = sql.Open("mysql", dsn) //链接数据库
	if err != nil {
		return err
	}
	err = Db.Ping() //检测数据库连接是否正常
	if err != nil {
		return err
	}
	return nil
}
