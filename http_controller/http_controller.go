package http_controller

import (
  "net/http"
  "fmt"
  "html/template"
  "database/sql"
)

func Start(db *sql.DB, stmt_information *sql.Stmt, stmt_accounts *sql.Stmt) {
  StartServer()
}

func StartServer() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    handler := Router(r, w)
    handler(r, w)
  })
  http.ListenAndServe(":8080", nil)
}



func Router(r *http.Request, w http.ResponseWriter) AnyHandler {
  return http.HandleFunc("POST /task", PrivateRequest(HandleCreateTask))
}
/* func router (db *sql.DB, stmt_information *sql.Stmt, stmt_accounts *sql.Stmt, w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    switch {
    case r.URL.Path == "/testcookieget":
      cookie := getCookieHandler(w, r)
      fmt.Fprintf(w, "cookie value: %v", cookie)
    case r.URL.Path == "/testcookieset":
      setCookieHandler(w, "text")
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
      isLoginEmpty(w, r)
    case r.URL.Path == "/register":
      renderTemplate(w, "register.html")
    case r.URL.Path == "/":
      http.Redirect(w, r, "/login", http.StatusFound)
    default:
      fmt.Fprintf(w, "404")
    }
  case http.MethodPost:
    switch {
    case r.URL.Path == "/login":
      isLoginEmpty(w, r)
    }
  default:
    fmt.Fprintln(w, "This method doesn't support")
  }
}

*/

func renderTemplate(w http.ResponseWriter, tmpl string) {
  t, err := template.ParseFiles("templates/" + tmpl)
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
