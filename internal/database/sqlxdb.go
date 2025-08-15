package database

import "github.com/jmoiron/sqlx"

var sqldb *sqlx.DB

func init() {
	var err error
	sqldb, err = sqlx.Connect("sqlite3", "gorm.db")

	if err != nil {
		panic("SQLX数据库连接失败, error=" + err.Error())
	}
	sqldb.SetMaxOpenConns(50)
	sqldb.SetMaxIdleConns(10)
}

func GetSqlxDB() *sqlx.DB {
	return sqldb
}
