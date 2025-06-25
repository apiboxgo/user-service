package user

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"user-service/api_init"
)

func GetItems(filterDto *RequestFilterUserDto) (*ResultListDTO, error) {

	query := api_init.GetDbh().Model(&User{})

	where := "(created_at, id) %s (?, ?)"
	orderCreatedAt, orderCreatedExists := filterDto.Orders["created_at"]

	if orderCreatedExists && strings.ToUpper(orderCreatedAt) != "DESC" {
		query.Order("created_at ASC")
		query.Order("id ASC")
		where = fmt.Sprintf(where, ">")
	} else {
		query.Order("created_at DESC")
		query.Order("id DESC")
		where = fmt.Sprintf(where, "<")
	}

	if filterDto.Cursor != "" && filterDto.LastTimestamp != "" {
		query.Where(where, filterDto.LastTimestamp, filterDto.Cursor)
	}

	if len(filterDto.Emails) > 0 {

		var likeConditions []string
		var likeArgs []interface{}
		for _, email := range filterDto.Emails {
			likeConditions = append(likeConditions, "email LIKE ?")
			likeArgs = append(likeArgs, email+"%")
		}
		query = query.Where("("+strings.Join(likeConditions, " OR ")+")", likeArgs...)
	}

	if filterDto.Limit <= 0 {
		filterDto.Limit = 10
	}

	var total int64
	var result []UserItemResultDto
	err := query.Count(&total).Error

	if err != nil {
		return nil, err
	}

	query.Limit(filterDto.Limit)
	query.Find(&result)

	lastItem := result[len(result)-1]
	resultDto := ResultListDTO{
		List:          result,
		Cursor:        lastItem.ID,
		LastTimestamp: lastItem.CreatedAt,
		Total:         total,
	}

	return &resultDto, nil
}

func GetOneByEmailAndPassword(email string, password string) (*UserItemResultDto, error) {

	if email == "" || password == "" {
		return nil, nil
	}

	result := User{}
	err := api_init.GetDbh().Raw("SELECT * FROM users WHERE email = $1 LIMIT 1", email).Scan(&result).Error

	if err == nil || err == sql.ErrNoRows {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &UserItemResultDto{
		ID:        result.ID,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}, nil
}

func GetOneById(requestUserIdDTO RequestUserIdDTO) (*UserItemResultDto, error) {

	if requestUserIdDTO.ID == "" {
		return nil, nil
	}

	result := UserItemResultDto{}
	err := api_init.GetDbh().Raw("SELECT * FROM users WHERE id = $1 LIMIT 1", requestUserIdDTO.ID).Scan(&result).Error
	return &result, err
}

func CreateUserItem(User User) (bool, error) {
	err := api_init.GetDbh().Create(&User).Error
	return result(err)
}

func PutUserItem(requestUserIdDTO RequestUserIdDTO, user map[string]interface{}) (bool, error) {
	err := api_init.GetDbh().Model(&User{}).Where("id = ?", requestUserIdDTO.ID).Updates(user).Error
	return result(err)
}

func PatchUserItem(User User) (bool, error) {
	err := api_init.GetDbh().Updates(&User).Error
	return result(err)
}

func DeleteUserItemById(id uuid.UUID) (bool, error) {
	err := api_init.GetDbh().Delete(&User{}, "id = ?", id.String()).Error
	return result(err)
}

func result(err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return true, nil
}
