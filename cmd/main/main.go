package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/gateway"
)

func main() {
	dir := ""
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		fmt.Println("Specify the yaml file path")
	}
	config, err := gatewayconfig.ReadConfig(dir)
	if err != nil {
		log.Fatal(err.Error())
	}

	handler := gateway.CreateGateway(config, 5000)

	http.ListenAndServe("0.0.0.0:5000", handler)

}
