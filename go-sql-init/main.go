package main

import (
	"database/sql"
	"fmt"
)

type SQL interface {
	Connect() (*sql.DB, error)
}

type MySQL struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (m *MySQL) Connect() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}


func main() {

}
