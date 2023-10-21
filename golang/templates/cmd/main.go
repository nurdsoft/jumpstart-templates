package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	// Version is the version of the application.
	VERSION string
	// COMMIT is the commit hash of the application build.
	COMMIT string
	// BUILDTIME is the time the application was built.
	BUILDTIME string
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(map[string]string{
			"version":   VERSION,
			"commit":    COMMIT,
			"buildtime": BUILDTIME,
		})
		w.Write(b)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
