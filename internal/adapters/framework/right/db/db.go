package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type DBAdapter struct {
	db *sql.DB
}

func NewDBAdapter(driverName, DBSourceName string) (*DBAdapter, error) {
	db, err := sql.Open(driverName, DBSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	return &DBAdapter{db: db}, nil
}

func (dbAdapter *DBAdapter) CloseConnection() {
	dbAdapter.db.Close()
}

func (dbAdapter *DBAdapter) SaveAnswerToHistory(answer int32, operation string) error {
	now := time.Now()
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(now, answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = dbAdapter.db.Exec(queryString, args...)
	if err != nil {
		log.Printf("Error inserting into history: %v", err)

		return err
	}
	return nil
}
