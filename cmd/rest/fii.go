package rest

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/go-funds-htmx/cmd/service"
)

func setupFiiEndpoints(e *echo.Echo, svc service.FiiService) {

	e.GET("/fii", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.POST("/fii/load", func(c echo.Context) error {

		fiiName := c.Request().PostFormValue("fii-name")
		log.WithField("fiiName", fiiName)
		funds, err := svc.GetYieldTickersFromFII(fiiName)

		if err != nil {
			log.WithError(err).Error("error coletando dados com fiiName")
			return c.Render(400, "index", err)
		}
		log.Info("dados coletados com sucesso")
		return c.Render(200, "fii/item", funds)
	})

	e.POST("/fii/load-tr", func(c echo.Context) error {
		fiiName := c.Request().PostFormValue("fii-name")
		funds, err := svc.GetYieldTickersFromFII(fiiName)

		if err != nil {
			log.WithError(err).Error("error coletando dados com fiiName")
			return c.Render(400, "index", err)
		}
		log.Info("dados coletados com sucesso")
		return c.Render(200, "fii/tr", funds)
	})
}

func setupApiFiiEndpoints(e *echo.Echo, svc service.FiiService) {
	e.GET("/api/fii", func(c echo.Context) error {
		fiiName := c.Request().URL.Query().Get("name")

		result, err := svc.GetDadosHistoricosFII(fiiName)

		if err != nil {
			return c.JSON(500, err)
		}

		return c.JSON(200, result)
	})
}