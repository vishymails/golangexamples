package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "db2admin"
	hostname = "127.0.0.1:3306"
	dbname   = "ibmdb"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error opening DB %s \n ", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error creating DB %s \n ", err)
		return
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error fetching DB %s \n ", err)
		return
	}

	log.Printf("rows effected %d \n ", no)

	db.Close()

	db, err1 := sql.Open("mysql", dsn(dbname))
	if err1 != nil {
		log.Printf("Error opening DB %s \n ", err1)
		return
	}
	defer db.Close()

	ctx1, cancelfunc1 := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc1()

	err = db.PingContext(ctx1)
	if err != nil {
		log.Printf("Error pinging DB %s \n ", err)
		return
	}

	log.Printf("Connected to %s DB Successfully", dbname)
}
