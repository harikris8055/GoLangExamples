//go program to check api health
package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			structs.Response

		}

	})

	http.ListenAndServe("localhost:3000", mux)

}
