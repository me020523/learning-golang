package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func helloHandler(resp http.ResponseWriter, request *http.Request) {
	store := sessions.NewCookieStore([]byte("123456"))
	session, err := store.New(request, "mysession")
	if err != nil {
		resp.WriteHeader(http.StatusBadGateway)
		resp.Write([]byte(err.Error()))
		return
	}
	session.Values["hello"] = "world"
	store.Save(request, resp, session)
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("hello,world"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
