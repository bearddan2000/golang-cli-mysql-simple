package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

const (
  host     = "db"
  port     = 3306
  user     = "maria"
  password = "pass"
  dbname   = "beverage"
)

func main() {
  mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
     user, password, host, port, dbname)

  db, err := sql.Open("mysql", mysqlInfo)

  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  db.Close()
}
