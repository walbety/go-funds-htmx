package canonical

type FundoImob struct {
	Name    string
	Ticker  string
	Dados   []RendimentoTicker
	Valores ValorTicker
}

type RendimentoTicker struct {
	Ticker     string  `json:"ticker"`
	Rendimento float32 `json:"rendimento"`
	Yield      float32 `json:"yield"`
	Data       string  `json:"data"`
}

type ValorTicker struct {
	Ticker    string      `json:"ticker"`
	DataValor []DataValor `json:"dataValor"`
}

type DataValor struct {
	Data  string
	Valor float32
}

type FiiValoresHistoricos struct {
	Valores []float32 `json:"valores"`
}
