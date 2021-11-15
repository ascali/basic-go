package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func OrderHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	r := c.Request()
	return c.Render(http.StatusOK, "order.html", map[string]interface{}{
		"name":       "Order",
		"msg":        "Hello Saya Dari Niomic",
		"Id_menu":    r.URL.Query()["Id_menu"][0],
		"Nama_menu":  r.URL.Query()["Nama_menu"][0],
		"Url_gambar": r.URL.Query()["Url_gambar"][0],
		"Harga":      r.URL.Query()["Harga"][0],
	})
}
