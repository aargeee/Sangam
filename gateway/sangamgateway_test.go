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

	var config = gatewayconfig.Config{
		Port: 5000,
		Paths: map[string]gatewayconfig.Path{
			"/hi": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: server.URL,
						},
					},
				},
			},
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

func TestGatewayMultipleRouting(t *testing.T) {

	HI_ROUTE := "/hi"
	HELLO_ROUTE := "/hello"

	handler := http.NewServeMux()
	handler.HandleFunc(HI_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "TEST STRING")
	})
	handler.HandleFunc(HELLO_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "NEW TEST STRING")
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	var config = gatewayconfig.Config{
		Port: 5000,
		Paths: map[string]gatewayconfig.Path{
			"/hi": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: server.URL,
						},
					},
				},
			},
			"/hello": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: server.URL,
						},
					},
				},
			},
		},
	}

	gw := gateway.CreateGateway(&config, 5000)
	t.Run("Test /hi route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, HI_ROUTE, nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		gw.ServeHTTP(res, req)
		assert.Equal[int](t, res.Code, http.StatusOK)
		assert.Equal[string](t, res.Body.String(), "TEST STRING")
	})

	t.Run("Test /hello route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, HELLO_ROUTE, nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		gw.ServeHTTP(res, req)
		assert.Equal[int](t, res.Code, http.StatusAccepted)
		assert.Equal[string](t, res.Body.String(), "NEW TEST STRING")
	})

}

func TestGatewayHeadersRelay(t *testing.T) {

	HI_ROUTE := "/hi"

	handler := http.NewServeMux()
	handler.HandleFunc(HI_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("RGHeader", "aargeee")
		fmt.Fprint(w, "TEST STRING")
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	var config = gatewayconfig.Config{
		Port: 5000,
		Paths: map[string]gatewayconfig.Path{
			"/hi": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: server.URL,
						},
					},
				},
			},
		},
	}

	gw := gateway.CreateGateway(&config, 5000)
	req, err := http.NewRequest(http.MethodGet, HI_ROUTE, nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	gw.ServeHTTP(res, req)
	assert.Equal[int](t, res.Code, http.StatusOK)
	assert.Equal[string](t, res.Body.String(), "TEST STRING")
	assert.Equal[string](t, res.Header().Get("RGHeader"), "aargeee")
}

func TestGatewayMultipleServers(t *testing.T) {

	HI_ROUTE := "/hi"
	HELLO_ROUTE := "/hello"

	handler := http.NewServeMux()
	handler.HandleFunc(HI_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "TEST STRING")
	})
	handler.HandleFunc(HELLO_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "http://%s", r.Host)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	nserver := httptest.NewServer(handler)
	defer nserver.Close()

	var config = gatewayconfig.Config{
		Port: 5000,
		Paths: map[string]gatewayconfig.Path{
			"/hi": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: server.URL,
						},
					},
				},
			},
			"/hello": {
				Methods: map[string]gatewayconfig.Method{
					"get": {
						Backend: gatewayconfig.Backend{
							Address: nserver.URL,
						},
					},
				},
			},
		},
	}

	gw := gateway.CreateGateway(&config, 5000)
	t.Run("Test /hi route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, HI_ROUTE, nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		gw.ServeHTTP(res, req)
		assert.Equal[int](t, res.Code, http.StatusOK)
		assert.Equal[string](t, res.Body.String(), "TEST STRING")
	})

	t.Run("Test /hello route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, HELLO_ROUTE, nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		gw.ServeHTTP(res, req)
		assert.Equal[int](t, res.Code, http.StatusAccepted)
		assert.Equal[string](t, res.Body.String(), nserver.URL)
	})

}
