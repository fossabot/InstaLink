package base

import (
	"net/http"

	"github.com/labstack/echo"
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
