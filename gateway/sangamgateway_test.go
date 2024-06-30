package gateway_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/constants"
	"github.com/aargeee/sangam/gateway"
	"github.com/alecthomas/assert/v2"
)

func TestSangamGateway(t *testing.T) {
	gw := gateway.CreateGateway(nil, 5000)
	req, err := http.NewRequest(http.MethodGet, constants.SANGAM_HEALTHZ, nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	gw.ServeHTTP(res, req)
	assert.Equal[int](t, res.Code, http.StatusOK)
}

func TestGatewayRouting(t *testing.T) {

	HI_ROUTE := "/hi"

	handler := http.NewServeMux()
	handler.HandleFunc(HI_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "TEST STRING")
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	var config = gatewayconfig.GatewayConfig{
		PORT: 5000,
		RoutesMap: map[string]string{
			HI_ROUTE: server.URL,
		},
	}

	gw := gateway.CreateGateway(&config, 5000)
	req, err := http.NewRequest(http.MethodGet, HI_ROUTE, nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	gw.ServeHTTP(res, req)
	assert.Equal[int](t, res.Code, http.StatusOK)
	assert.Equal[string](t, res.Body.String(), "TEST STRING")
}
