package mockserver

import (
	"fmt"
	"net/http"
)

func EchoHandler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("PATH: %s\n", r.URL.Path)
	})
	return router
}
