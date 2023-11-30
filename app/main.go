package main

import (
	"log"

	"github.com/hyperversalblocks/mailer/pkg/api"
)

func main() {
	if err := api.Init(); err != nil {
		log.Fatal(err)
	}
}
