package http_controller

import (
  "net/http"
  "fmt"
  "html/template"
  "database/sql"
  )

func Start() {
  StartServer()
}

func DbAndStmt(database *sql.DB, statement *sql.Stmt) {
 db = database
 stmt = statement
}

func StartServer() {
     http.HandleFunc("/", router)
     http.ListenAndServe(":8080", nil)
}

func router (w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/html; charset=utf-8")

 switch r.Method {
  case http.MethodGet:
   switch {
   case r.URL.Path == "/list":
    renderTemplate(w, "list.html")
    List(w)
   case r.URL.Path == "/create":
    renderTemplate(w, "create.html")
     Create(r.URL.Query(), w)
   case r.URL.Path == "/clear": 
    renderTemplate(w, "clear.html")
    Clear(w)
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
