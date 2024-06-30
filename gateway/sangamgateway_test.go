package gateway_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/gateway"
	"github.com/alecthomas/assert/v2"
)

var config = gatewayconfig.GatewayConfig{}

func TestSangamGateway(t *testing.T) {
	gw := gateway.CreateGateway(&config, 5000)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	gw.ServeHTTP(res, req)
}
