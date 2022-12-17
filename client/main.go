package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	resp, err := http.Get("http://go-server:8080/")
	if err != nil {
		return c.String(http.StatusInternalServerError, "get response error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "read response error")
	}

	return c.String(http.StatusOK, string(body))
}
