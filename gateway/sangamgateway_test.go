package gateway_test

import (
	"testing"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/gateway"
)

var config = gatewayconfig.GatewayConfig{}

func TestSangamGateway(t *testing.T) {
	gateway.CreateGateway(config, 5000)
}
