package main

import (
  "test/http_controller"
  "fmt"
 _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "os"
 )


func main() {
  db, stmt, err := open_file("mydatabase.db")
    if err != nil {
    return
  }
  defer db.Close()
  defer stmt.Close()
  
  if os.Args[1] == "server" {
  http_controller.Start(db, stmt)
}
}



func open_file(filename string) (*sql.DB, *sql.Stmt, error) {
  var db *sql.DB
  var err error
  db, err = sql.Open("sqlite3", filename)
 if err != nil {
  return nil, nil, err
 }
 _, err = db.Exec(`
  CREATE TABLE IF NOT EXISTS information (
  info TEXT,
  time TEXT)
  `)
 if err != nil {
  fmt.Println(err)
  return nil, nil, err
 }
 var stmt *sql.Stmt
 stmt, err = db.Prepare("INSERT INTO information(info, time) VALUES(?, ?)")
 if err != nil {
  return nil, nil, err
 }

 return db, stmt, nil
}
