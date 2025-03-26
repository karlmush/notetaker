package http_controller

import (
  "crypto/rand"
  "net/http"
  "fmt"
)


func setCookieHandler(w http.ResponseWriter, session_id string) {
  cookie := http.Cookie{
  Name:     "SessionIdCookie",
  Value:    session_id,
  } 
  http.SetCookie(w, &cookie)
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) (cookieHash string){

  cookie, err := r.Cookie("SessionIdCookie")
  if err != nil {
    fmt.Fprint(w, "Error while read cookie\n")
    return
  }
  return cookie.Value
}

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
