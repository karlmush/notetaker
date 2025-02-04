package main

import (
  "test/http_controller"
  "fmt"
 _ "github.com/mattn/go-sqlite3"
  "database/sql"
 )

var db *sql.DB
var stmt *sql.Stmt

func main() {
  err := open_file("mydatabase.db")
    if err != nil {
    return
  }
  defer db.Close()
  defer stmt.Close()
  http_controller.DbAndStmt(db, stmt)
  http_controller.Start()
}



func open_file(filename string) error {
  var err error
  db, err = sql.Open("sqlite3", filename)
 if err != nil {
  return err
 }
 _, err = db.Exec(`
  CREATE TABLE IF NOT EXISTS information (
  info TEXT,
  time TEXT)
  `)
 if err != nil {
  fmt.Println(err)
  return err
 }
 stmt, err = db.Prepare("INSERT INTO information(info, time) VALUES(?, ?)")
 if err != nil {
  return err
 }

 return nil
}
