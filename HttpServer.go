package main

import (
	"net/http"
)

func httpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("Request received...")
		println(r.Method, "(", r.Proto, ") =>", r.RequestURI)
		w.Write([]byte("Hi, there!"))
	})
	println("Server is listnening at http://localhost:3000")

	http.ListenAndServe("localhost:3000", mux)
}

/*
Open http://localhost:3000
or
Use `curl localhost:3000` to send request to port 3000
    `curl -i localhost:3000` for request details
    `curl -X POST localhost:3000` for spacifice method of request
*/
