package main

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func main() {
	gzip := gziphandler.GzipHandler(http.HandlerFunc(Handle))
	http.Handle("/", gzip)

	http.ListenAndServe(":8080", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world :" + r.URL.Path))
}
