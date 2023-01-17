package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 8000, "端口")

func main() {
	flag.Parse()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("hello mc,port:%d", *port)))
	})
	log.Println("server listen at:", *port)
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), nil)
}
