package http_controller

import (
  "fmt"
  "net/http"
  "html/template"
)

    type LoginData struct {
     Login string
     LoginError string
     PasswordError string
    }

func isLoginEmpty(w http.ResponseWriter, r *http.Request) bool {
  if r.Method == http.MethodGet {
    renderLoginTemplate(w, LoginData{})
  } else if r.Method == http.MethodPost {
      login := r.FormValue("login")
      password := r.FormValue("password")
      data := LoginData{Login: login}

      if login == "" {
        data.LoginError = "Login required"
      }
      if password == "" {
        data.PasswordError = "Password required"
      }
      if data.LoginError != "" || data.PasswordError != ""{
        renderLoginTemplate(w, data)
      }
      return false
    }
  return true
}




    func renderLoginTemplate(w http.ResponseWriter, data LoginData) {
     tmpl, err := template.ParseFiles("templates/login.html")
     if err != nil {
       fmt.Fprint(w, "Error while opening login.html")
       return
     }
      err = tmpl.Execute(w, data)
     if err != nil {
       fmt.Fprint(w, "Error while executing data to login.html")
       return
     }
}
