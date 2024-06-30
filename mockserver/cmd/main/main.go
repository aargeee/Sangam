package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aargeee/sangam/mockserver"
)

func main() {
	port := 3000
	var err error
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	fmt.Printf("EchoServer running on localhost:%d...\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mockserver.EchoHandler())
}
