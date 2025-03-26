package http_controller

import (
  "net/http"
  "database/sql"
  "html/template"
  "fmt"
)

type AnyHandler = func(r *http.Request, w http.ResponseWriter)

func getPublicContext (_ *http.Request) (error, *PublicContext) {
  return nil, &PublicContext{
    DB: nil, // сюда db
  }
} 

type PublicContext struct {
  DB *sql.DB
}

type PublicHandler = func (*http.Request, http.ResponseWriter, *PublicContext)

func PublicRequest(handler PublicHandler) AnyHandler {
  return func(r *http.Request, w http.ResponseWriter) {
    err, ctx := getPublicContext(r)
    if err != nil {
      //unauthorized 401 
      fmt.Println(err)
    }
    handler(r, w, ctx)
  }
}

func getPrivateContext(r *http.Request) (error, *PrivateContext) {
  err, ctx := getPublicContext(r)
  if err != nil {
    return err, nil
  }
  return nil, &PrivateContext{
    PublicContext: *ctx,
    UserID: "123",
  }
}


type  PrivateContext struct {
  PublicContext
  UserID string
}

type PrivateHandler = func(*http.Request, http.ResponseWriter, *PrivateContext)

func PrivateRequest(handler PrivateHandler) AnyHandler {
  return func(r *http.Request, w http.ResponseWriter) {
    err, ctx := getPrivateContext(r)
    if err != nil {
      //unauthorized 401
      fmt.Println(err)
    }
    handler(r, w, ctx)
  }
}
/*
func Router(path string) AnyHandler {
  switch path{
  case "/private":
    return PrivateRequest(CreatePrivateHandler)
  default:
    return PublicRequest(NotFoundHandler)
  }
}
*/
func HandleCreateTask (r *http.Request, w http.ResponseWriter, ctx *PrivateContext) {
  t, err := template.ParseFiles("templates/create.html")
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}


func NotFoundHandler(r *http.Request, w http.ResponseWriter, ctx *PublicContext) {
  http.NotFound(w, r)
  fmt.Fprint(w, "NOT FOUND")
}


