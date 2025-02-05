package http_controller

import(
  "database/sql"
  "time"
  "net/url"
  "net/http"
  "fmt"
)


func Create(params url.Values, w http.ResponseWriter, stmt *sql.Stmt){
 TimeNow := time.Now()
 date := TimeNow.Format("2006-02-01")
 time := TimeNow.Format("15:04:05")
 date_and_time := date + " at " + time

 information := params.Get("info")

 if information == "" {
 fmt.Fprintf(w, "Empty input")
 return
 } else {
 _, err := stmt.Exec(information, date_and_time)
 if err != nil {
 fmt.Fprintf(w, "Exec error")
 return
 }

 fmt.Fprintf(w, "Success")
 return
 }
}
