package main

import (
	"fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main()  {
  r := mux.NewRouter()
  r.HandleFunc("/hello", handler).Methods("GET")
  r.HandleFunc("/portfolio", portfolio).Methods("GET")
  http.ListenAndServe(":7996", r)
}

func handler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Himmath na haarna, saahas na chodna!")
}

func portfolio(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Portfolio will be here, coming soon.")
}
