package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"otherpay-test/common"
)

var DB *sql.DB

func MysqlClient() *sql.DB {
	if DB == nil {
		config := common.SerConf.MysqlCfg
		InitDB(config)
	}
	return DB
}


func InitDB(config common.MysqlConfig) (err error) {
	dsn := config.Dsn
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}