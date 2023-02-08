package postgresql

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
    _ "github.com/lib/pq"
)

type DbConnect struct {
    Db *sql.DB
}

func (d *DbConnect) Connect() {
    host := os.Getenv("POSTGRES_HOST")
    port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")


    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    d.Db = db
}

func (d *DbConnect) GetCur() *sql.DB {
    return d.Db
}