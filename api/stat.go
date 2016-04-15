package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetStatInfo provides statical information of traffic
func GetStatInfo(c *echo.Context) error {
	// TODO: get statistic informaiton
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}
