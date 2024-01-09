package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
