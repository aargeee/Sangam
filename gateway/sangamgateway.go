package gateway

type Gateway struct {
	port int
}

func CreateGateway(port int) *Gateway {
	return &Gateway{
		port: port,
	}
}

func (g *Gateway) ListenAndServe() {

}
