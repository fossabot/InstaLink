
import (
	"github.com/labstack/echo"
)

// Init is a router initializing function
func InitRouter() *echo.Echo {

	e := echo.New()

	e.GET("/", status.RootView)

	// Routings
	u1 := e.Group("/utility/api/v1")
	{
		UtilityAPIRoutings(u1)
	}

	return e
}
