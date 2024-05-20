package service

import "github.com/walbety/go-funds-htmx/cmd/canonical"

type FiiService interface {
	GetYieldTickersFromFII(Fundo string) ([]canonical.RendimentoTicker, error)
	GetDadosHistoricosFII(Fundo string) (canonical.FiiValoresHistoricos, error)
	// GetValuesHistoric(name string, dataInicio time.Time, dataFim time.Time) (canonical.ValorTicker, error)
}