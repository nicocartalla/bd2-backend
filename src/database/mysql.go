package database

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)

var (
	dbUri         string
)

func init() {
	config, err := config.LoadConfig("./")
	if err != nil {
		utils.ErrorLogger.Fatal("cannot load config:", err)
	}
	dbUri = config.DBUri

}

func dBInit() *sql.DB {
	var err error
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}

func QueryDB(query string) (*sql.Rows, error) {
	db := dBInit()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func QueryRowDB(query string, parameter any) (*sql.Row, error) {
	db := dBInit()
	defer db.Close()
	row := db.QueryRow(query, parameter)
	return row, nil
}

func InsertDB(insert string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(insert)
	if err != nil {
		return 0, err
	}
	return d.LastInsertId()
}

func InsertDBParams(insert string, params ...any) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(insert, params...)
	if err != nil {
		return 0, err
	}
	return d.LastInsertId()
}

func DeleteDB(delete string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(delete)
	if err != nil {
		return 0, err
	}
	return d.RowsAffected()
}

func UpdateDB(update string) (int64, error) {
	db := dBInit()
	defer db.Close()
	d, err := db.Exec(update)
	if err != nil {
		return 0, err
	}
	return d.RowsAffected()
}
