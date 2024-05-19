package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"ecurrency/pkg"

	"github.com/gin-gonic/gin"
)

// The following shoud be divided into service & repository parts

func (h *Handler) getRates(c *gin.Context) {
	link := "https://api.monobank.ua/bank/currency"
	USD_Code := 840
	UA_Code := 980

	response, err := http.Get(link)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	body, err := io.ReadAll(response.Body)

	if err != nil {
		c.Error(err)
	}

	var rates []pkg.Rate

	json.Unmarshal(body, &rates)

	var output pkg.Rate
	for _, rate := range rates {
		if int(rate.CurrencyCodeA) == USD_Code && int(rate.CurrencyCodeB) == UA_Code {
			output = rate
		}
	}

	c.JSON(http.StatusOK, output.RateBuy)
}

func (h *Handler) getEmail(c *gin.Context) {
	email := c.PostForm("email")

	c.JSON(http.StatusOK, gin.H{
		"email": email,
	})
}
