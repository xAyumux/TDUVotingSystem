package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/xAyumux/TDUVotingSystem/model"
)

func Connect() *sql.DB {
	errLoad := godotenv.Load("db.env")
	if errLoad != nil {
		panic("Error loading .env file")
	}
	dsn := os.Getenv("DSN")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Select() []model.Polls {
	results := []model.Polls{}

	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description FROM polls")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	polls := model.Polls{}

	for rows.Next() {
		err := rows.Scan(&polls.Id, &polls.Title, &polls.Description)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, polls)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func SelectQuery(id int64) []model.Polls {
	results := []model.Polls{}

	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description FROM polls WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	polls := model.Polls{}

	for rows.Next() {
		err := rows.Scan(&polls.Id, &polls.Title, &polls.Description)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, polls)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func Insert(userid string, id, result int64) int64 {
	db := Connect()
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO votes(userid, id, result) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	insertResult, err := stmtInsert.Exec(userid, id, result)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertID
}
