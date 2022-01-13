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

func SelectPoll() (model.Polls, int) {
	results := model.Polls{}

	db := Connect()
	defer db.Close()

	// idが一番大きいデータを取得
	// 並び替え後、最初の1行を取得
	rows, err := db.Query("SELECT id, title, description FROM polls ORDER BY id DESC LIMIT 1")
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
	}
	results = polls

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// idを使いvotesテーブルからresultが同じものの件数を取得
	rowsVotes, err := db.Query("SELECT id, result, COUNT(result) FROM votes WHERE id = ? GROUP BY result", polls.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsVotes.Close()

	votes := model.Votes{}
	var number int = 0

	for rowsVotes.Next() {
		err := rowsVotes.Scan(&votes.Id, &votes.Result, number)
		if err != nil {
			log.Fatal(err)
		}
	}

	return results, number
}

func SelectQuery(id string) (model.Polls, int) {
	results := model.Polls{}

	db := Connect()
	defer db.Close()

	// 指定されたidからデータを取得
	// rows, err := db.Query("SELECT id, title, description FROM polls WHERE id = ?", id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	polls := model.Polls{}

	// for rows.Next() {
	// 	err := rows.Scan(&polls.Id, &polls.Title, &polls.Description)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	err := db.QueryRow("SELECT id, title, description FROM polls WHERE id = ?", id).Scan(&polls.Id, &polls.Title, &polls.Description)
	if err != nil {
		log.Fatal(err)
	}
	results = polls

	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// idを使いvotesテーブルからresultが同じものの件数を取得
	rowsVotes, err := db.Query("SELECT id, result, COUNT(result) FROM votes WHERE id = ? GROUP BY result", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsVotes.Close()

	votes := model.Votes{}
	var number int = 0

	for rowsVotes.Next() {
		err := rowsVotes.Scan(&votes.Id, &votes.Result, number)
		if err != nil {
			log.Fatal(err)
		}
	}

	return results, number
}

func SelectPollsList() []model.Polls {
	results := []model.Polls{}

	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description FROM polls")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	votes := model.Polls{}

	for rows.Next() {
		err := rows.Scan(&votes.Id, &votes.Title, &votes.Description)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, votes)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func InsertPolls(title, description string) int64 {
	db := Connect()
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO polls(title, description) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	insertResult, err := stmtInsert.Exec(title, description)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertID
}

func InsertVotes(userid, id, result string) int64 {
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
