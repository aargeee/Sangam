package gateway

import (
	"io"
	"net/http"

	gatewayconfig "github.com/aargeee/sangam/GatewayConfig"
	"github.com/aargeee/sangam/constants"
)

type Gateway struct {
	config *gatewayconfig.Config
	port   int
	http.Handler
}

func CreateGateway(config *gatewayconfig.Config, port int) *Gateway {

	router := http.NewServeMux()
	router.HandleFunc(constants.SANGAM_HEALTHZ, func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		baseURL := config.Paths[r.URL.Path].Methods["get"].Backend.Address
		res, err := http.Get(baseURL + r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		for key, values := range res.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

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
