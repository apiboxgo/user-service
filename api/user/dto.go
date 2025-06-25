package user

import (
	"github.com/google/uuid"
	"time"
)

// ============================== Request DTO ==========================================================================

type RequestFilterUserDto struct {
	Emails        []string          `form:"emails[]"`
	Limit         int               `form:"limit"`
	Cursor        string            `form:"cursor"`
	LastTimestamp string            `json:"lastTimestamp"`
	Orders        map[string]string `json:"orders[]"`
}

type RequestUserDTO struct {
	Email     string `form:"email" binding:"required,email" example:"Some user email"`
	Password  string `form:"password" binding:"required" example:"Some user password"`
	CreatedAt string `form:"created_at" example:"2022-01-01T00:00:00Z"`
	UpdatedAt string `form:"updated_at" example:"2022-01-01T00:00:00Z"`
	DeletedAt string `form:"deleted_at" example:"2022-01-01T00:00:00Z"`
}

type RequestUserIdDTO struct {
	ID string `uri:"id" binding:"required,uuid" example:"987fbc97-4bed-5078-9f07-9141ba07c9f3"`
}

type RequestUserByEmailAndPasswordDto struct {
	Email    string `form:"email" binding:"required,email"  example:"Some user email"`
	Password string `form:"password" binding:"required"  example:"Some user password"`
}

// ============================== Response DTO =========================================================================

type ErrorResponseDto struct {
	Message string `json:"message"`
}

type SuccessResponseDto struct {
	Message string `json:"message"`
}

type UserItemResultDto struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ResultListDTO struct {
	List          []UserItemResultDto `json:"list"`
	Cursor        uuid.UUID           `json:"cursor"`
	LastTimestamp time.Time           `json:"lastTimestamp"`
	Total         int64               `json:"total"`
}
