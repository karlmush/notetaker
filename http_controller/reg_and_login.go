package http_controller

import (
  "fmt"
  "net/http"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request, db *sql.DB, stmt *sql.Stmt) {
  err := r.ParseForm()
  if err != nil {
    fmt.Fprint(w, "Can't parse form\n")
    return
  }
  password := r.FormValue("password")
  login := r.FormValue("login")
  code, _ := isUser_checkpass(db, login, password)

  if len(r.Form) == 2 { // if 2 forms = login(only 1 form for login and 1 for pass)
    if code == 2 {
    _, session := isUser_checkpass(db, login, password)
    if session != ""{
    setCookieHandler(w, session)
      fmt.Fprint(w, "Success log in\n")
    } else {
      fmt.Fprint(w, "Empty session id\n")
    }
  } else if code == 1 {
    fmt.Fprint(w, "Found user. Invalid password\n")
  } else if code == 3 {
    fmt.Fprintf(w, "Not all inputs filled\n")
  } else {
    fmt.Fprintf(w, "Can't find user with this login. Check it or go to registration page\n")
  }
  if len(r.Form) == 3  { // if 3 forms = register, not login(bc of 2 forms for pass)

    login := r.FormValue("login")
    password := r.FormValue("password_1")
    if r.FormValue("password_1") != r.FormValue("password_2") {
      password = ""
      fmt.Fprintf(w, "Passwords doesn't match\n")
    }
    code, _ := isUser_checkpass(db, login, password)
    if code == 1 || code == 2 {
      fmt.Fprintf(w, "This login is busy\n")
    } else if code == 3{
      fmt.Fprintf(w, "Not all inputs filled\n")
    } else {
      session_id, err := gencookie()
      if err != nil {
        fmt.Fprint(w, "Error while creating session_id for cookies\n")
        return
      }
      password_hash, err := HashPassword(password)
      if err != nil {
        fmt.Fprint(w, "Error while hashing password\n")
      }
      registration(w, stmt, login, password_hash, session_id)
      fmt.Fprint(w, "Succes registration")
    }
  }
} 
}

func isUser_checkpass(db *sql.DB, login string, password string)(int, string) {

  if login == "" || password == "" {
    return 3, ""
  } else {

    rows, err := db.Query("SELECT * FROM accounts")
    if err != nil {
      panic(err)
    }
    defer rows.Close()

    for rows.Next() {
      var login_db string
      var password_hash_db string
      var session string
      err := rows.Scan(&login_db, &password_hash_db, &session)
      if err != nil {
        panic(err)
      }
      if login == login_db {
        if CheckPasswordHash(password, password_hash_db){
          return 2, session
        }
        return 1, ""
      }
    }
  }
  return 0, ""
}

func registration(w http.ResponseWriter, stmt *sql.Stmt, login string, password string, session_id string) {


  if login == "" || password == "" {
    fmt.Fprintf(w, "Not all inputs filled\n")
    return
  } else {
    _, err := stmt.Exec(login, password, session_id)
    if err != nil {
      fmt.Fprintf(w, "Exec error\n")
      return
    } else {
      fmt.Fprintf(w, "Successfully registration\n")
    }
  }
}


func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
