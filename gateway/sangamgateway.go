package gateway

import (
	"net/http"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
)

type Gateway struct {
	config gatewayconfig.GatewayConfig
	port   int
	http.Handler
}

func CreateGateway(config gatewayconfig.GatewayConfig, port int) *Gateway {

	router := http.NewServeMux()

	return &Gateway{
		config:  config,
		port:    port,
		Handler: router,
	}
}
