package http_controller 

import(
  "net/http"
  "fmt"
  "database/sql"
)

func Clear(w http.ResponseWriter, db *sql.DB) {
 _, err := db.Exec("DELETE FROM information")
 if err != nil {
  fmt.Fprintf(w, "Database error: %v", err)
  return
 }
 fmt.Fprintf(w, "Succesfully")
}
