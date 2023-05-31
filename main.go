package main

import (
	"gotest/container"
	"log"
	"net/http"
)

func main() {

	server := &http.Server{Addr: ":8088", Handler: container.GetmainCon()}
	log.Fatal(server.ListenAndServe())

}
