package gateway

import (
	"net/http"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
)

type Gateway struct {
	config gatewayconfig.GatewayConfig
	port   int
}

func CreateGateway(config gatewayconfig.GatewayConfig, port int) *Gateway {
	return &Gateway{
		config: config,
		port:   port,
	}
}

func (g *Gateway) GetHandler() (*http.Handler, error) {
	return nil, nil
}
