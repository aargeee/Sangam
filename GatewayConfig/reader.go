package gatewayconfig

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfig(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
