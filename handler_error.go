package main

import "net/http"

func HandlerError(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, 500, "something went wrong")
}
