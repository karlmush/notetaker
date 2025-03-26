
package main

import (
    "errors"
    "log"
    "net/http"
    "fmt"
)

func main() {
    // Start a web server with the two endpoints.
    mux := http.NewServeMux()
    mux.HandleFunc("/set", setCookieHandler)
    mux.HandleFunc("/get", getCookieHandler)

    log.Print("Listening...")
    err := http.ListenAndServe(":3000", mux)
    if err != nil {
        log.Fatal(err)
    }
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookie := http.Cookie{
        Name:     "exampleCookie",
        Value:    "Hello world!",
        Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }

    // Use the http.SetCookie() function to send the cookie to the client.
    // Behind the scenes this adds a `Set-Cookie` header to the response
    // containing the necessary cookie data.
    http.SetCookie(w, &cookie)

    // Write a HTTP response as normal.
    fmt.Fprint(w, "Cookie set")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("exampleCookie")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Error(w, "cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        return
    }

    w.Write([]byte(cookie.Value))
}

