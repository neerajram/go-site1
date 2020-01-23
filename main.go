package main

import (
	"fmt"
  "net/http"
)

func main()  {
  http.HandleFunc("/", handler)
  http.HandleFunc("/resume", resume)
  http.ListenAndServe(":7996", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Written in go!")
}

func resume(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Resume will be here, coming soon.")
}
