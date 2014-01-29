package main

import (
	"net/http"
	"regexp"
	"fmt"
)

func adHandler(w http.ResponseWriter, r *http.Request, fid string) {
	if r.Method == "POST" {
		fmt.Printf("r.Method = %s\n", r.Method)
//		w.Header().Set("Content-Type", "image/jpeg")
		w.Write([]byte("This is GET request " + r.Method))
	} else {
		http.NotFound(w, r)
		return
	}
}

var validPath = regexp.MustCompile("^/(fid)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", makeHandler(adHandler))
	http.ListenAndServe(":8080", nil)
}
