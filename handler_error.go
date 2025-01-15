package main

import (
	"net/http"

	"github.com/Komakivan/go-scraper/json_res"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	json_res.ResponseError(w, 500, "something went wrong")
}
