package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/walbety/go-funds-htmx/cmd/canonical"
)

const (
	URL_INFOMONEY_HISTORICO         = "https://fii-api.infomoney.com.br/api/v1/fii/provento/historico?Ticker="
	URL_INFOMONEY_VALORES           = "https://fii-api.infomoney.com.br/api/v1/fii/cotacao/historico/grafico"
	URL_INFOMONEY_HISTORICO_GRAFICO = "https://fii-api.infomoney.com.br/api/v1/fii/cotacao/historico/grafico?Ticker=%s&DataInicio=%s&DataFim=%s"
)

type InfomoneySvc struct {
}

func NewInfomoneyService() FiiService {
	return InfomoneySvc{}
}

func (s InfomoneySvc) GetYieldTickersFromFII(Fundo string) ([]canonical.RendimentoTicker, error) {

	url := fmt.Sprint(URL_INFOMONEY_HISTORICO, Fundo)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error at Get")
		return []canonical.RendimentoTicker{}, err
	}

	var rendimentos []canonical.RendimentoTicker

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error at readall")
		return []canonical.RendimentoTicker{}, err
	}

	err = json.Unmarshal(body, &rendimentos)
	if err != nil {
		fmt.Println("error at Unmarshal")
		return []canonical.RendimentoTicker{}, err
	}

	return rendimentos, nil
}

func (s InfomoneySvc) GetDadosHistoricosFII(Fundo string) (canonical.FiiValoresHistoricos, error) {

	agora := time.Now()
	url := fmt.Sprintf(URL_INFOMONEY_HISTORICO_GRAFICO, Fundo,
		agora.AddDate(-1, 0, 0).Format("02-01-2006"), agora.Format("02-01-2006"))

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error at Get")
		return canonical.FiiValoresHistoricos{}, err
	}

	var rendimentos canonical.ValorTicker

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error at readall")
		return canonical.FiiValoresHistoricos{}, err
	}

	err = json.Unmarshal(body, &rendimentos)
	if err != nil {
		fmt.Println("error at Unmarshal")
		return canonical.FiiValoresHistoricos{}, err
	}

	var result canonical.FiiValoresHistoricos

	for _, val := range rendimentos.DataValor {
		result.Valores = append(result.Valores, val.Valor)
	}

	fmt.Println(agora.Format("2006-01-02"))

	return result, nil
}
