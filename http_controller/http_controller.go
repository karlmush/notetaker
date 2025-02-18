package http_controller

import (
  "net/http"
  "fmt"
  "html/template"
  "database/sql"
  )

func Start(db *sql.DB, stmt_information *sql.Stmt, stmt_accounts *sql.Stmt) {
  StartServer(db, stmt_information, stmt_accounts)
}

func StartServer(db *sql.DB, stmt_information *sql.Stmt, stmt_accounts *sql.Stmt) {
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      router(db, stmt_information, stmt_accounts, w, r)
     })
     http.ListenAndServe(":8080", nil)
}

func router (db *sql.DB, stmt_information *sql.Stmt, stmt_accounts *sql.Stmt, w http.ResponseWriter, r *http.Request) {
 switch r.Method {
  case http.MethodGet:
   switch {
   case r.URL.Path == "/testcookiegen":
         getCookieHandler(w, r)
   case r.URL.Path == "/list":
    renderTemplate(w, "list.html")
    List(w, db)
   case r.URL.Path == "/create":
    renderTemplate(w, "create.html")
     Create(r.URL.Query(), w, stmt_information)
   case r.URL.Path == "/clear": 
    renderTemplate(w, "clear.html")
    Clear(w, db)
   case r.URL.Path == "/main":
     renderTemplate(w, "main.html")
   case r.URL.Path == "/login":
    renderTemplate(w, "login.html")
   case r.URL.Path == "/register":
    renderTemplate(w, "register.html")
   default:
    renderTemplate(w, "login.html")
  }
  case http.MethodPost:
     login(w, r, db, stmt_accounts)
  default:
    fmt.Fprintln(w, "This method doesn't support")
   }
 }


func renderTemplate(w http.ResponseWriter, tmpl string) {
  t, err := template.ParseFiles("templates/" + tmpl)
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
