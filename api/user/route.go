package user

import (
	"github.com/gin-gonic/gin"
)

const UriUser = "/user"
const UriUserGetByEmail = "/get-by-email"
const UriUserGetById = "/:id"
const UriUserGetByIdS = "/%s"

func InitUserRoutes(route *gin.Engine) {
	group := route.Group(UriUser)
	route.GET(UriUser, GetUsersListByFilter)
	group.GET(UriUserGetById, GetUserById)
	route.POST(UriUser, CreateUser)
	group.POST(UriUserGetByEmail, GetUserByEmail)
	group.PUT(UriUserGetById, PutUserItemById)
	group.PATCH(UriUserGetById, PatchUserById)
	group.DELETE(UriUserGetById, DeleteUserById)
}
