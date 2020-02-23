package main

import (
	"fmt"
	"net/http"
  "github.com/gomodule/redigo/redis"
)

var (
	c redis.Conn
	err error
	reply interface{}
)

func main() {
/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})
*/

  // Handle pooling using several redis clients... //

	c, err = redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Print(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	  s, _ := redis.String(c.Do("GET", "zway/kjokken/temperatur"))
		fmt.Fprintf(w, "Hello World! I'm a HTTP server and the temperature is %s Celsius!",s)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serving webserver on localhost:3001")
	http.ListenAndServe(":3001", nil)
}


func addPipe() {
	fmt.Println("Will eventually add a new pipe.")
}
