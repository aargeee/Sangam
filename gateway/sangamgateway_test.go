package gateway_test

import (
	"testing"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/gateway"
	"github.com/alecthomas/assert/v2"
)

var config = gatewayconfig.GatewayConfig{}

func TestSangamGateway(t *testing.T) {
	ms_gateway := gateway.CreateGateway(config, 5000)
	assert.NoError(t, ms_gateway.ListenAndServe())
}
