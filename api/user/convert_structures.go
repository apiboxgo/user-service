package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
	"user-service/dictionary"
	"user-service/utils"
)

func convertRequestUserDTOToUser(c *gin.Context, requestUserIdDTO RequestUserIdDTO, requestUserDTO RequestUserDTO, user *User) {
	id, err := uuid.Parse(requestUserIdDTO.ID)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	user.ID = id
	createdAt, err := time.Parse(time.RFC3339, requestUserDTO.CreatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	user.CreatedAt = createdAt

	updatedAt, err := time.Parse(time.RFC3339, requestUserDTO.UpdatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	user.UpdatedAt = updatedAt

	deletedAt, err := time.Parse(time.RFC3339, requestUserDTO.DeletedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	user.DeletedAt = deletedAt
}

func convertRequestUserDTOToMap(c *gin.Context, requestUserDTO RequestUserDTO) map[string]interface{} {

	result := map[string]interface{}{}

	createdAt, err := time.Parse(time.RFC3339, requestUserDTO.CreatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if createdAt.IsZero() {
		createdAt = time.Time{}
	}

	result["created_at"] = createdAt
	updatedAt, err := time.Parse(time.RFC3339, requestUserDTO.UpdatedAt)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if updatedAt.IsZero() {
		updatedAt = time.Time{}
	}

	result["updated_at"] = updatedAt

	deletedAt, err := time.Parse(time.RFC3339, requestUserDTO.DeletedAt)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if deletedAt.IsZero() {
		deletedAt = time.Time{}
	}

	result["deleted_at"] = deletedAt

	return result
}
