package main

import (
  "test/http_controller"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "os"
  "github.com/joho/godotenv"
  _ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  databaseURL := os.Getenv("TURSO_DATABASE_URL")
  apiKey := os.Getenv("TURSO_AUTH_TOKEN")

  url := fmt.Sprintf("%v?authToken=%v", databaseURL, apiKey)

  db, stmt_information, stmt_accounts, err := open_db(url)
  if err != nil {
    return
  }
  defer db.Close()

  defer stmt_information.Close()
  defer stmt_accounts.Close()

  if os.Args[1] == "server" {
    http_controller.Start(db, stmt_information, stmt_accounts)
  } else if os.Args[1] != "server" {
    fmt.Println("usage: go run main.go server")
  }
}


func open_db(url  string) (*sql.DB, *sql.Stmt, *sql.Stmt, error) {
  var db *sql.DB
  var err error
  db, err = sql.Open("libsql", url)
  if err != nil {
    return nil, nil, nil, err
  }
  _, err = db.Exec(`
  CREATE TABLE IF NOT EXISTS information (
    info TEXT,
    time TEXT,
    session TEXT)
    `)
    if err != nil {
      fmt.Println(err)
      return nil, nil, nil, err
    }

    _, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS accounts (
      login TEXT,
      password TEXT,
      session TEXT)
      `)
      if err != nil {
        fmt.Println(err)
        return nil, nil, nil, err
      }

      var stmt_information *sql.Stmt
      stmt_information, err = db.Prepare("INSERT INTO information(info, time) VALUES(?, ?)")
      if err != nil {
        return nil, nil, nil, err
      }

      var stmt_accounts *sql.Stmt
      stmt_accounts, err = db.Prepare("INSERT INTO accounts(login, password, session) VALUES(?, ?, ?)")
      if err != nil {
        return nil, nil, nil, err
      }

      return db, stmt_information, stmt_accounts, nil
    }
