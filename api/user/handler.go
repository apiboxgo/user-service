package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
	"user-service/dictionary"
	_ "user-service/docs"
	"user-service/utils"
)

// ================================== Get user by Email and Password ===================================================

//	@title			Getting user by Email and Password
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetUserByEmailAndPassword @Summary Getting user by Email and Password
// @Description Getting user by Email and Password
// @Tags user
// @Accept json
// @Produce json
// @Param        request body RequestUserByEmailAndPasswordDto true "Sent data"
// @Success 200 {array} RequestUserDTO
// @Router /user/get-by-email-and-password [post]
func GetUserByEmailAndPassword(c *gin.Context) {

	var requestUserByEmailAndPasswordDto RequestUserByEmailAndPasswordDto

	if err := c.ShouldBindJSON(&requestUserByEmailAndPasswordDto); err != nil {
		utils.Dump(err)
		utils.LogError(dictionary.ErrorParsingRequestBody, err)
		c.JSON(http.StatusUnprocessableEntity, &ErrorResponseDto{
			Message: err.Error(),
		})
	}
	resultDto, err := GetOneByEmailAndPassword(requestUserByEmailAndPasswordDto.Email, requestUserByEmailAndPasswordDto.Password)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	if resultDto == nil || resultDto.ID == uuid.Nil {
		message := fmt.Sprintf(dictionary.UserNotFound, requestUserByEmailAndPasswordDto.Email)
		c.JSON(http.StatusNotFound, &ErrorResponseDto{
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, resultDto)
}

// ================================== Get user by ID ===================================================================

//	@title			Getting user by id
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetUserById @Summary Getting user by id
// @Description Getting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User id (UUID)"
// @Success 200 {array} RequestUserDTO
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) {

	requestDto, _ := parseDtoId(c)

	if requestDto.ID == "" {
		return
	}

	resultDto, err := GetOneById(requestDto)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	if resultDto.ID == uuid.Nil {
		message := fmt.Sprintf(dictionary.UserByIdNotFound, requestDto.ID)
		c.JSON(http.StatusNotFound, &ErrorResponseDto{
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, resultDto)
}

// ================================== Get Users by filter ==============================================================

//	@title			Getting Users
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetUsersListByFilter @Summary Getting Users
// @Description Getting Users
// @Tags Users
// @Accept json
// @Produce json
// @Param names query []string false "Name"
// @Param limit query int false "Limit"
// @Param cursor query string false "cursor (las id uuid)"
// @Param lastTimestamp query string false "lastTimestamp"
// @Param orders[created_at] query []string false "Filter created_at Like min-max (example: 2025-06-11T08:28:51.400404Z)"
// @Success 200 {array} RequestUserDTO
// @Router /user [get]
func GetUsersListByFilter(c *gin.Context) {
	emails := c.QueryArray("emails")
	lastTimestamp := c.Query("lastTimestamp")
	ordersCreatedAt := c.Query("orders[created_at]")
	var orders map[string]string

	if ordersCreatedAt != "" {
		orders = map[string]string{
			"created_at": strings.ToUpper(ordersCreatedAt),
		}
	} else {
		orders = map[string]string{
			"created_at": "DESC",
		}
	}

	if lastTimestamp != "" {
		_, err := time.Parse(time.RFC3339Nano, lastTimestamp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(dictionary.ErrorParsingFilter, "lastTimestamp", lastTimestamp)})
			return
		}
	}

	if len(emails) > 0 {
		emails = strings.Split(emails[0], ",")
	}

	requestFilterUserDto := &RequestFilterUserDto{
		Emails:        emails,
		LastTimestamp: lastTimestamp,
		Orders:        orders,
	}

	if err := c.ShouldBindQuery(&requestFilterUserDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultListDTO, err := GetItems(requestFilterUserDto)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	listLength := len(resultListDTO.List)

	if listLength == 0 {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusNotFound, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})

		return
	}
	c.JSON(http.StatusOK, resultListDTO)
}

// ================================== Delete user by ID ================================================================

//	@title			Deleting user by id
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// DeleteUserById @Summary Getting user by id
// @Description Deleting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User id (UUID)"
// @Success 200 {array} RequestUserDTO
// @Router /user/{id} [delete]
func DeleteUserById(c *gin.Context) {
	_, id := parseDtoId(c)

	if id == uuid.Nil {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	isDeleted, err := DeleteUserItemById(id)
	if err != nil || !isDeleted {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.UserDeletedSuccessful,
	})
}

// ================================== Patch user by ID =================================================================
//	@title			Patch user
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// PatchUserById     godoc
// @Summary      Patch user
// @Description  Update only sent fields
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User id (UUID)"
// @Param        request body RequestUserDTO true "Updated data"
// @Success      200 {object}  UserItemResultDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /user/{id} [patch]
func PatchUserById(c *gin.Context) {
	_, id := parseDtoId(c)
	User, _ := parseRequestBody(c)
	User.ID = id

	isUpdated, err := PatchUserItem(User)

	if err != nil || !isUpdated {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// ================================== Put user by ID ===================================================================
//	@title			Put user
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// PutUserItemById   godoc
// @Summary      Put user
// @Description  Update all sent fields
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path string true "User id (UUID)"
// @Param        request body  RequestUserDTO true "Updated data"
// @Success      200 {object}  SuccessResponseDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /user/{id} [put]
func PutUserItemById(c *gin.Context) {
	requestIdDto, _ := parseDtoId(c)
	_, requestUserPostDTO := parseRequestBody(c)
	UserMap := convertRequestUserDTOToMap(c, requestUserPostDTO)
	UserMap["updated_at"] = time.Now()
	isUpdated, err := PutUserItem(requestIdDto, UserMap)

	if err != nil || !isUpdated {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// ================================== Create (post) user ===============================================================
//	@title			Create user
//	@version		1.0
//	@description	Create user
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// CreateUser godoc
// @Summary      Create user
// @Description  Create user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body RequestUserDTO true "Sent data"
// @Success      200 {object}  SuccessResponseDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /user [post]
func CreateUser(c *gin.Context) {

	User, _ := parseRequestBody(c)
	isCreated, err := CreateUserItem(User)

	if err != nil || !isCreated {
		utils.Dump(err)
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusCreated, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// === Sys

func parseDtoId(c *gin.Context) (RequestUserIdDTO, uuid.UUID) {
	var requestUserIdDTO RequestUserIdDTO
	var id uuid.UUID
	if err := c.ShouldBindUri(&requestUserIdDTO); err != nil {
		utils.LogError(dictionary.ErrorParsingRequestBody, err)
		c.JSON(http.StatusUnprocessableEntity, &ErrorResponseDto{
			Message: err.Error(),
		})
	} else {
		id, err = uuid.Parse(requestUserIdDTO.ID)
		if err != nil {
			utils.LogError(dictionary.UserByIdNotFound, err)
			c.JSON(http.StatusNotFound, &ErrorResponseDto{
				Message: dictionary.UserByIdNotFound,
			})
		}
	}

	return requestUserIdDTO, id
}

func parseRequestBody(c *gin.Context) (User, RequestUserDTO) {
	var requestUserPostDTO RequestUserDTO
	if err := c.ShouldBindJSON(&requestUserPostDTO); err != nil {
		utils.LogError(dictionary.ErrorParsingRequestBody, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
	}

	if requestUserPostDTO.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(requestUserPostDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.LogError(dictionary.ErrorParsingRequestBody, err)
			c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
				Message: dictionary.SomethingWrong,
			})
		}
		requestUserPostDTO.Password = string(hash)

	}

	return User{
			Email:    requestUserPostDTO.Email,
			Password: requestUserPostDTO.Password,
		},
		requestUserPostDTO
}
