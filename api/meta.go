package api

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/westlab/door-api/model"
)

// CreateMeta creates meta data
func CreateMeta(c echo.Context) error {
	// TODO: Create Meta
	return c.JSON(http.StatusCreated, "{'hello': 'world'}")
}

// GetMeta gets meta data
func GetMeta(c echo.Context) error {
	// TODO: Get Meta
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// GetMetaByName get a meta data by name
func GetMetaByName(c echo.Context) error {
	// TODO: Get Meta

	name := c.Param("name")
	m := model.SelectSingleMeta(name)
	if m == nil {
		return c.JSONBlob(http.StatusOK, []byte("{}"))
		//return c.JSON(http.StatusOK, )
	}
	return c.JSON(http.StatusOK, m)
}

// UpdateMeta updates meta data
func UpdateMeta(c echo.Context) error {
	// TODO: Update  Meta
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}

// DeleteMeta deletes meta data
func DeleteMeta(c echo.Context) error {
	// TODO: Delete Meta
	return c.JSON(http.StatusOK, "{'hello': 'world'}")
}
