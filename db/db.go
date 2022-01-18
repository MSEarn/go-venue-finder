package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MSEarn/go-venue-finder/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewPool(dbCfg *config.Mysql) (*sql.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
	)

	dbPool, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dbPool.SetMaxOpenConns(int(dbCfg.MaxOpenConns))
	dbPool.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetime) * time.Second)
	fmt.Printf("Database %s starting... \n", dbCfg.DBName)

	return dbPool, nil
}
