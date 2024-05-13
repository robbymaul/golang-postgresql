package testing

import (
	"context"
	"log"
	"testing"

	"github.com/robbymaul/golang-postgresql.git/connection"
	"github.com/robbymaul/golang-postgresql.git/model"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()
	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()

	tx, _ := conn.Begin(ctx)

	var user model.User
	user.Username = "robby"
	user.Password = "robby"

	sql := "INSERT INTO users (username, password) VALUES ($1,$2)"

	_, err = tx.Exec(ctx, sql, user.Username, user.Password)
	if err != nil {
		tx.Rollback(ctx)
		log.Fatal(err.Error())
	}

	tx.Commit(ctx)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	conn, _ := connection.GetConnection()
	defer conn.Close()

	user := model.User{
		Username: "robby",
		Password: "ganteng",
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	sql := "UPDATE users SET password=$1 where username=$2"

	_, err = tx.Exec(ctx, sql, user.Password, user.Username)
	if err != nil {
		tx.Rollback(ctx)
		log.Fatal(err.Error())
	}

	tx.Commit(ctx)
}

func TestSelect(t *testing.T) {
	var user model.User

	ctx := context.Background()
	conn, _ := connection.GetConnection()

	sql := "SELECT username, password FROM users where username=$1"

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := tx.Query(ctx, sql, "robby")
	if err != nil {
		log.Fatal(err.Error())
		tx.Rollback(ctx)
	}

	for rows.Next() {
		rows.Scan(&user.Username, &user.Password)
	}

	assert.Equal(t, "robby", user.Username)
	assert.Equal(t, "ganteng", user.Password)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	conn, _ := connection.GetConnection()

	tx, _ := conn.Begin(ctx)

	sql := "DELETE FROM users WHERE username = $1"

	_, err := tx.Exec(ctx, sql, "robby")
	if err != nil {
		tx.Rollback(ctx)
		log.Fatal(err.Error())
	}

	tx.Commit(ctx)

	assert.Nil(t, err)

}
