package storage

import "database/sql"

var DB sql.DB

func init() {
	//todo: connect database
}

func New() *sql.DB {
	// todo: return db variable
	return &DB
}

func close() {
	//todo: close db connection
	DB.Close()
}
