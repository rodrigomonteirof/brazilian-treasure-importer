package main

import (
	"log"

	"github.com/rodrigomonteirof/brazilian-treasure-importer/config"
	"github.com/rodrigomonteirof/brazilian-treasure-importer/http"
	"github.com/rodrigomonteirof/brazilian-treasure-importer/tesouro"
)

func main() {
	csvUrl, err := tesouro.GetCSVUrl(config.TesouroDiretoAPIUrl())
	if err != nil {
		log.Fatalf("failed to get CSV URL: %v", err)
	}

	err = http.Download(csvUrl, config.QuotesCSVPath())
	if err != nil {
		log.Fatalf("failed to download quotes: %v", err)
	}
}
