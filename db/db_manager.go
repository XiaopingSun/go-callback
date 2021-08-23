package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"qiniu/setting"
)

var DB *sql.DB

func ConnectMysql() (err error) {
	// 连接数据库
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		setting.Setting().Database.User,
		setting.Setting().Database.Password,
		setting.Setting().Database.Network,
		setting.Setting().Database.Host,
		setting.Setting().Database.Port,
		setting.Setting().Database.DBName)
	DB, err = sql.Open(setting.Setting().Database.Type, conn)
	return
}

func CreateTable(sql string) (error) {
	_, err := DB.Exec(sql)
	return err
}

func Insert(sql string, args ...interface{}) (int64, error) {
	result, err := DB.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	lastInsertID,err := result.LastInsertId()
	return lastInsertID, err
}