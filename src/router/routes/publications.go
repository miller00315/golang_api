package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Function:               controllers.ListPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
}
