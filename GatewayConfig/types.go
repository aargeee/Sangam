package gatewayconfig

type Backend struct {
	Address string `yaml:"address"`
}

type Method struct {
	Backend Backend `yaml:"x-backend"`
}

type Path struct {
	Methods map[string]Method `yaml:",inline"`
}

type Config struct {
	Port  int             `yaml:"PORT"`
	Paths map[string]Path `yaml:"paths"`
}
