package http_controller

import (
  "fmt"
  "net/http"
  "database/sql"
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
    _, hash := isUser_checkpass(db, login, password)
    setCookieHandler(w, hash)
    fmt.Fprintf(w, "Login and password matched\n")
  } else if code == 1 {
    fmt.Fprintf(w, "Found user. Invalid password\n")
  } else if code == 3 {
    fmt.Fprintf(w, "Not all inputs filled\n")
  } else {
    fmt.Fprintf(w, "Can't find user with this login. Check it or go to registration page\n")
  }
 } else if len(r.Form) == 3  { // if 3 forms = register, not login(bc of 2 forms for pass)
   login := r.FormValue("login")
   var password string
   if r.FormValue("password_1") == r.FormValue("password_2\n") {
     password = r.FormValue("password_1\n")
   } else {
    fmt.Fprintf(w, "Passwords doesn't match\n")
    return
   }
   if code == 1 || code == 2 {
    fmt.Fprintf(w, "This login is busy\n")
   } else if code == 3{
    fmt.Fprintf(w, "Not all inputs filled\n")
  } else {
    hash, err := gencookie()
    if err != nil {
    fmt.Fprint(w, "Error while creating hash for cookies\n")
    return
    }
    registration(w, stmt, login, password, hash)
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
    var password_db string
    var hash string
    err := rows.Scan(&login_db, &password_db, &hash)
    if err != nil {
      panic(err)
    }
    if login == login_db {
      if password == password_db {
        return 2, hash
      }
     return 1, ""
    }
   }
  }
  return 0, ""
}

func registration(w http.ResponseWriter, stmt *sql.Stmt, login string, password string, hash string) {


  if login == "" || password == "" {
    fmt.Fprintf(w, "Not all inputs filled\n")
    return
  } else {
   _, err := stmt.Exec(login, password, hash)
   if err != nil {
   fmt.Fprintf(w, "Exec error\n")
   return
 } else {
   fmt.Fprintf(w, "Successfully registration\n")
 }
  }
}
