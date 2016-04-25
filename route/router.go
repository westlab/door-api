package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/westlab/door-api/api"
	"github.com/westlab/door-api/conf"
)

// Init initialize door api
func Init(c conf.Config) *echo.Echo {
	e := echo.New()

	// Debug
	e.SetDebug(c.AppDebug)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	v1 := e.Group("/v1")
	{
		// browsing
		v1.Get("/browsings/:id", api.GetBrowsingByID)
		v1.Get("/browsings", api.GetBrowsings)
		v1.Post("/browsings", api.CreateBrowsing)
		v1.Put("/browsings/:id", api.UpdateBrowsing)
		v1.Delete("/browsings/:id", api.DeleteBrowsing)

		v1.Get("/browsing_histogram", api.GetBrowsingHistorgram)
		v1.Get("/ip_rank", api.GetIPRank)
		v1.Get("/domain_rank", api.GetDomainRank)
		v1.Get("/word_rank", api.GetWordRank)

		v1.Get("/stat_info", api.GetStatInfo)
	}

	return e
}
