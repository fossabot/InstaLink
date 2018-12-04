package base

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shotastage/instalink/core"
)

func RootView(c echo.Context) error {
	return c.HTML(http.StatusOK, `
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>InstaLink</title>
</head>
<body>
	<h1>InstaLink Server</h1>
	<p>
	This is API server, please connect and manage via RESTAPI.
	</p>
</body>
</html>
	`)
}

func ServiceReady(c echo.Context) error {

	m := &core.Message{
		Message: "Ready to provide all services.",
		Type:    "String",
	}

	return c.JSONPretty(http.StatusOK, m, " ")
}

func PingPong(c echo.Context) error {

	m := &core.Message{
		Message: "Pong",
		Type:    "String",
	}

	return c.JSONPretty(http.StatusOK, m, " ")
}
