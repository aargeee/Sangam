package gateway_test

import (
	"testing"

	"github.com/aargeee/sangam/gateway"
)

func TestSangamGateway(t *testing.T) {
	ms_gateway := gateway.CreateGateway(5000)
	ms_gateway.ListenAndServe()
}
