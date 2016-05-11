package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/westlab/door-api/model"
)

// CreateBrowsing cretes browsing record
func CreateBrowsing(c echo.Context) error {
	// TODO: implement create browsing function
	return c.JSON(http.StatusCreated, "{'hello': 'world'}")
}

// UpdateBrowsing updates browsing record
func UpdateBrowsing(c echo.Context) error {
	// TODO: implement update browsing function
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetBrowsings get and search browsing record
func GetBrowsings(c echo.Context) error {
	q := c.QueryParam("q")
	size, _ := strconv.Atoi(c.QueryParam("size"))
	return c.JSON(http.StatusOK, model.GetBrowsings(q, int64(size)))
}

// DeleteBrowsing delete browsing record
func DeleteBrowsing(c echo.Context) error {
	// TODO: implement delete browsing function
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetBrowsingHistogram returns browsing histogram in specific time window
// with fin grain manner.
func GetBrowsingHistogram(c echo.Context) error {
	window, _ := strconv.Atoi(c.Param("window"))
	duration, _ := strconv.Atoi(c.Param("duration"))
	if window == 0 {
		window = 10
	}
	if duration == 0 {
		duration = 300
	}
	return c.JSON(http.StatusOK,
		model.GetBrowsingHistogram(int64(duration), int64(window)))
}

// GetBrowsingBySrcIP gets browsing by src ip
func GetBrowsingBySrcIP(c echo.Context) error {
	// call model.GetBrowsingByID
	// example: how to get URL or GET params
	srcIP := c.Param("src_ip")
	return c.JSON(http.StatusOK, model.GetBrowsingBySrcIP(srcIP))
}
