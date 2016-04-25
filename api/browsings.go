package api

import (
	"github.com/labstack/echo"
	"net/http"
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

// GetBrowsing get and search browsing record
func GetBrowsings(c echo.Context) error {
	// TODO: implement get browsing function
	// call model.GetBrowsings
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// DeleteBrowsing delete browsing record
func DeleteBrowsing(c echo.Context) error {
	// TODO: implement delete browsing function
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetBrowsingHistorgram get browsing histogram in specific time window
// with fin grain manner.
func GetBrowsingHistorgram(c echo.Context) error {
	// TODO: implement get browsing histogram
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

func GetBrowsingByID(c echo.Context) error {
	// call model.GetBrowsingByID
	// example: how to get URL or GET params
	// id := c.Param("id")
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}
