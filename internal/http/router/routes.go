package router

import (
	"GOLANG/internal/http/handler"

	"github.com/labstack/echo/v4"
)

// membuat struct route
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// membuat fungsi untuk mengembalikan route
// pada func ini perlu login krna private
func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
	}
}

// membuat fungsi untuk mengembalikan route
// pada func ini tdk perlu login krna public
func PrivateRoutes(UserHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: UserHandler.CreateUser,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: UserHandler.GetAllUser,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: UserHandler.UpdateUser,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: UserHandler.GetUserByID,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: UserHandler.DeleteUser,
		},
	}
}

//NOTE :
//MENGAPA TERDAPAT 2 FUNC DIATAS? YAITU PUBLIC DAN PRIVATE
//KAREN DI SERVER.GO KITA BUAT GROUP API, DAN KITA MEMBAGI ROUTE YANG PERLU LOGIN DAN TIDAK PERLU LOGIN
// YAITU PUBLIC DAN PRIVATE

//note ;
//untuk menjalankan nya setelah port 8080 ditambahin /api/v1
// karna di server.go kita membuat group API
