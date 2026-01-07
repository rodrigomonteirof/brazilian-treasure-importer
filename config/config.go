package config

import (
	"fmt"
	"time"
)

func QuotesCSVPath() string {
	return fmt.Sprintf("/tmp/quotes-%s.csv", time.Now().Format("2006-01-02"))
}

func TesouroDiretoAPIUrl() string {
	return "https://www.tesourotransparente.gov.br/ckan/api/3/action/package_show?id=taxas-dos-titulos-ofertados-pelo-tesouro-direto"
}
