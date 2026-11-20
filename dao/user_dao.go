package dao

import (
	"context"
	"database/sql"
	"fmt"
	"kaizen/model"
	"log"
	"os"

	"github.com/oklog/ulid"
)

var db *sql.DB

func OpenDB() *sql.DB {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_USER_PWD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	if mysqlUser == "" || mysqlUserPwd == "" || mysqlDatabase == "" {
		log.Fatal("MYSQL_USER、MYSQL_USER_PWD、MYSQL_DATABASE is empty")
	}

	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
	return db
}

func SearchUsersByName(name string) (error, []model.User) {
	users := make([]model.User, 0)
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return err, users
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return err, users
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return err, users
	}
	return nil, users
}

func InsertUser(context context.Context, id ulid.ULID, req model.User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context, "INSERT INTO user (id, name, age) VALUES (?, ?, ?)", id.String(), req.Name, req.Age)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
