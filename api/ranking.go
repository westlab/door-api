package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/westlab/door-api/model"
)

// GetIPRank provides ranking of frequent SrcIP in specific time window
func GetIPRank(c echo.Context) error {
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetWordRank provides ranking of frequent word in specific time window
func GetWordRank(c echo.Context) error {
	size, _ := strconv.Atoi(c.Param("size"))
	return c.JSON(http.StatusOK, model.GetWordCount(int64(size)))
}

// GetDomainRank provides ranking of frequent domain in specific time window
func GetDomainRank(c echo.Context) error {
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}
