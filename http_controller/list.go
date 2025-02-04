package http_controller

import (
  "fmt"
  "database/sql"
  "net/http"
)

var db *sql.DB

func List(w http.ResponseWriter) {
 rows, err := db.Query("SELECT * FROM information")
 if err != nil {
   panic(err)
 }
 defer rows.Close()

 for rows.Next() {
  var information_output string
  var time_output string
  err := rows.Scan(&information_output, &time_output)
  if err != nil {
    panic(err)
  }
  fmt.Fprintf(w, "%s, you added %s<br>", information_output, time_output)
 }
}

