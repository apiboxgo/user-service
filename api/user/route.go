package user

import (
	"github.com/gin-gonic/gin"
)

const UriUser = "/user"
const UriUserGetByEmailAndPassword = UriUser + "/get-by-email-and-password"
const UriUserGetById = UriUser + "/:id"
const UriUserGetByIdS = UriUser + "/%s"

func InitUserRoutes(route *gin.Engine) {
	route.GET(UriUser, GetUsersListByFilter)
	route.GET(UriUserGetById, GetUserById)
	route.POST(UriUser, CreateUser)
	route.POST(UriUserGetByEmailAndPassword, GetUserByEmailAndPassword)
	route.PUT(UriUserGetById, PutUserItemById)
	route.PATCH(UriUserGetById, PatchUserById)
	route.DELETE(UriUserGetById, DeleteUserById)
}
