package server

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Message":      "hello world",
		"StatusBanner": StatusBannerInformation(),
		"Services": map[string]interface{}{
			"Categories": ServicesInformation(),
		},
	})
}

func StartServer() {
	templates, err := template.ParseFS(files, "templates/*.html")
	if err != nil {
		fmt.Println("error")
	}

	t := &Template{
		templates: templates,
	}

	e := echo.New()
	e.Debug = true
	e.Renderer = t

	e.Static("/static", "assets")
	e.GET("/", Index)
	e.Start(":8000")
}
