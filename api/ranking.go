package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetIPRank provides ranking of frequent SrcIP in specific time window
func GetIPRank(c *echo.Context) error {
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetWordRank provides ranking of frequent word in specific time window
func GetWordRank(c *echo.Context) error {
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetDomainRank provides ranking of frequent domain in specific time window
func GetDomainRank(c *echo.Context) error {
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}
