package gateway

import (
	"net/http"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/constants"
)

type Gateway struct {
	config *gatewayconfig.GatewayConfig
	port   int
	http.Handler
}

func CreateGateway(config *gatewayconfig.GatewayConfig, port int) *Gateway {

	router := http.NewServeMux()
	router.HandleFunc(constants.SANGAM_HEALTHZ, func(w http.ResponseWriter, r *http.Request) {})

	return &Gateway{
		config:  config,
		port:    port,
		Handler: router,
	}
}
