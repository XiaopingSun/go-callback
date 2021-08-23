package callback

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"qiniu/db"
)

func CreateTable() (err error) {
	createTable := `CREATE TABLE IF NOT EXISTS callbacks (
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	time VARCHAR(64) NULL DEFAULT NULL,
	source VARCHAR(64) NULL DEFAULT NULL,
	remoteip VARCHAR(64) NULL DEFAULT NULL,
	header TEXT NULL DEFAULT NULL,
	body TEXT NULL DEFAULT NULL
	);`
	err = db.CreateTable(createTable)
	return err
}

func insertCallback(callback callbackItem) {
	insertCallback := "INSERT INTO callbacks(time, source, remoteip, header, body) value(?, ?, ?, ?, ?)"
	index, err := db.Insert(insertCallback, callback.time, callback.source, callback.remoteip, callback.requestHeader, callback.requestBody)
	if err != nil {
		fmt.Println("insert callback failed:", err)
	} else {
		fmt.Println("insert callback success, last index is ", index)
	}
}