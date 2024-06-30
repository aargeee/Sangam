package gateway

import (
	"io"
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
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		baseURL := config.RoutesMap[r.URL.Path]
		res, err := http.Get(baseURL + r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		w.WriteHeader(res.StatusCode)
		_, err = io.Copy(w, res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return &Gateway{
		config:  config,
		port:    port,
		Handler: router,
	}
}
