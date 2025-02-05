package http_controller

import (
  "net/http"
  "fmt"
  "html/template"
  "database/sql"
  )

func Start(db *sql.DB, stmt *sql.Stmt) {
  StartServer(db, stmt)
}

func StartServer(db *sql.DB, stmt *sql.Stmt) {
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      router(db, stmt, w, r)
     })
     http.ListenAndServe(":8080", nil)
}

func router (db *sql.DB, stmt *sql.Stmt, w http.ResponseWriter, r *http.Request) {
 switch r.Method {
  case http.MethodGet:
   switch {
   case r.URL.Path == "/list":
    renderTemplate(w, "list.html")
    List(w, db)
   case r.URL.Path == "/create":
    renderTemplate(w, "create.html")
     Create(r.URL.Query(), w, stmt)
   case r.URL.Path == "/clear": 
    renderTemplate(w, "clear.html")
    Clear(w, db)
   default:
    renderTemplate(w, "main_page.html")
  }
  default:
  renderTemplate(w, "main_page.html")
   }
 }


func renderTemplate(w http.ResponseWriter, tmpl string) {
  t, err := template.ParseFiles("templates/" + tmpl)
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
