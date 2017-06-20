package main

import (
	"net/http"

	"github.com/byuoitav/adcp-control-microservice/handlers"
	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/hateoas"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":8012"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/:address/volume/set/:level", handlers.SetVolume)
	secure.GET("/:address/power/on", handlers.PowerOn)
	secure.GET("/:address/power/standby", handlers.PowerStandby)
	secure.GET("/:address/volume/mute", handlers.Mute)
	secure.GET("/:address/volume/unmute", handlers.UnMute)

	//status endpoints
	secure.GET("/:address/volume/level", handlers.VolumeLevel)
	secure.GET("/:address/volume/mute/status", handlers.MuteStatus)
	secure.GET("/:address/power/status", handlers.PowerStatus)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
