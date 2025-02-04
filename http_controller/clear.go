package http_controller 

import(
  "net/http"
  "fmt"
)

func Clear(w http.ResponseWriter) {
 _, err := db.Exec("DELETE FROM information")
 if err != nil {
  fmt.Fprintf(w, "Database error: %v", err)
  return
 }
 fmt.Fprintf(w, "Succesfully")
}
