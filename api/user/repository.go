package user

import (
	"fmt"
	"github.com/apiboxgo/library-utils/api_init"
	"github.com/google/uuid"
	"strings"
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

func GetOneByEmail(email string) (*UserItemFullResultDto, error) {

	if email == "" {
		return nil, nil
	}

	var result UserItemFullResultDto
	err := api_init.GetDbh().Raw("SELECT * FROM users WHERE email = $1 LIMIT 1", email).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
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
