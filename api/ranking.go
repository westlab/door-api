package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/westlab/door-api/model"
)

// GetIPRank provides ranking of frequent SrcIP in specific time window
func GetIPRank(c echo.Context) error {
	duration, _ := strconv.Atoi(c.Param("duration"))
	r := model.GetBrowsingRank("src_ip", int64(duration))
	if r == nil {
		return c.JSONBlob(http.StatusOK, []byte("[]"))
	}

	return c.JSON(http.StatusOK, r)
}

// GetWordRank provides ranking of frequent word in specific time window
func GetWordRank(c echo.Context) error {
	size, _ := strconv.Atoi(c.Param("size"))
	r := model.GetWordCount(int64(size))
	if r == nil {
		return c.JSONBlob(http.StatusOK, []byte("[]"))
	}

	return c.JSON(http.StatusOK, r)
}

// GetDomainRank provides ranking of frequent domain in specific time window
func GetDomainRank(c echo.Context) error {
	duration, _ := strconv.Atoi(c.Param("duration"))
	r := model.GetBrowsingRank("domain", int64(duration))
	if r == nil {
		return c.JSONBlob(http.StatusOK, []byte("[]"))
	}
	return c.JSON(http.StatusOK, r)
}
