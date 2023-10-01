package main

import(
  "io"
  "log"
  "net/http"
)

func main(){
  
  helloHandler := func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, Scott!\n")
  }

  homeHandler := func(w http.ResponseWriter, req *http.Request){
    io.WriteString(w, "Home Page, now I need to connect a frontend to here.")
  }

  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/", homeHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))

}


