package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnFollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/following",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowing,
		RequiresAuthentication: true,
	},
}
