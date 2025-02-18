package http_controller

import (
  "crypto/rand"
  "net/http"
  "fmt"
)


func gencookie() (x string, err error) { 
  infoCookie, err :=  generateCookie(20)
  if err != nil{
  return "", err
  }
  if infoCookie == "" {
    return "", nil
  }
  return infoCookie, nil
}

func setCookieHandler(w http.ResponseWriter, hash string) {

  w.Header().Set("Content-Type", "text/html")

  cookie := http.Cookie{
  Name: "HashCookie",
  Value: hash,
  Path: "/",
  MaxAge: 3600,
  HttpOnly:  false,
  Secure: false,
} 

  fmt.Fprint(w, "Attempting to set cookie\n")
  http.SetCookie(w, &cookie)
  fmt.Fprint(w, "Cookie succesfully set\n")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {

  cookie, err := r.Cookie("HashCookie")
  if err != nil {
    fmt.Fprint(w, "Error while read cookie\n")
    return
  }
  fmt.Fprintf(w, "%v", cookie)
}

func generateCookie(length int) (string, error) {
  const charset = "abcdefghijklmnopqrstuvwxyz" +
      "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
      "0123456789"

  cookieFingerprint := make([]byte, length)
  _, err := rand.Read(cookieFingerprint)
  if err != nil {
    return "", err
  }

  for i := range cookieFingerprint{
   cookieFingerprint[i] = charset[int(cookieFingerprint[i])%len(charset)]
  }
  return string(cookieFingerprint), nil
}
