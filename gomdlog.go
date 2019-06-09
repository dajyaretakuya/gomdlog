package main

import (
	"log"
	"net/http"

	"github.com/dajyaretakuya/gomdlog/route"
)

func main() {
	route.Dispatch()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
