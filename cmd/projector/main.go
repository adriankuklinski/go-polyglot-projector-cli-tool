package main

import (
	"fmt"
	"log"

	"github.com/adriankuklinski/go-polyglot-projector-cli-tool/pkg/config"
)

func main() {
	opts, err := config.GetOpts()

	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	config, err := config.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	fmt.Printf("config: %+v", config)
}
