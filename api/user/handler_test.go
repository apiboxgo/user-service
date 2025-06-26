package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apiboxgo/library-utils/api_init"
	"github.com/apiboxgo/library-utils/dictionary"
	"github.com/apiboxgo/library-utils/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

var db *gorm.DB

func init() {
	api_init.TestInit("../../")
	db = api_init.InitGlobal.Dbh
}

func TestGetUsersList(t *testing.T) {
	clearDbTableUser(t)

	_, err := createUsers(14)

	u, err := url.Parse(UriUser)
	if err != nil {
		panic(err)
	}

	const limit = 3

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, limit, len(result.List))
	assert.Equal(t, "test_user_14@user.com", result.List[0].Email)
}

func TestGetUsersListWithFilterEmail(t *testing.T) {
	clearDbTableUser(t)

	_, err := createUsers(14)
	u, err := url.Parse(UriUser)
	if err != nil {
		panic(err)
	}

	const limit = 5

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("emails", "test_user_10@user.com,test_user_9@user.com,test_user_8@user.com")

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)
	utils.Dump(result.List)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(result.List))
	assert.Equal(t, "test_user_10@user.com", result.List[0].Email)
	assert.Equal(t, "test_user_9@user.com", result.List[1].Email)
	assert.Equal(t, "test_user_8@user.com", result.List[2].Email)
}

func TestGetUsersListPagination(t *testing.T) {
	clearDbTableUser(t)
	Users, err := createUsers(14)

	u, err := url.Parse(UriUser)
	if err != nil {
		panic(err)
	}

	const limit = 3

	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("cursor", Users[11].ID.String())
	q.Set("lastTimestamp", Users[11].CreatedAt.Format(time.RFC3339Nano))

	u.RawQuery = q.Encode()

	var result ResultListDTO
	w := sendRequest(t, u.String(), "GET", nil, &result)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, limit, len(result.List))
	assert.Equal(t, "test_user_11@user.com", result.List[0].Email)
}

func TestGetUserById_NotFoundResult(t *testing.T) {
	clearDbTableUser(t)
	fakeId := "987fbc97-4bed-5078-9f07-9141ba07c9f3"
	var result ErrorResponseDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, fakeId), "GET", nil, &result)

	assert.Equal(t, http.StatusNotFound, w.Code)

	message := fmt.Sprintf(dictionary.UserByIdNotFound, fakeId)
	assert.Equal(t, message, result.Message)
}

func TestGetUserById_WrongIdFormat(t *testing.T) {
	clearDbTableUser(t)
	fakeId := "987fbc97"
	var result ErrorResponseDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, fakeId), "GET", nil, &result)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	assert.Equal(t, "Key: 'RequestUserIdDTO.ID' Error:Field validation for 'ID' failed on the 'uuid' tag", result.Message)
}

func TestGetUserById_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)
	User := User{
		Email:    "test_user_1@user.com",
		Password: "123123",
	}
	if err := db.Create(&User).Error; err != nil {
		t.Fatal(err)
	}

	var result UserItemResultDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, User.ID.String()), "GET", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, User.ID, result.ID)
}

func TestGetUserByEmail_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)
	User := User{
		Email:    "test_user_1@user.com",
		Password: "123123",
	}
	if err := db.Create(&User).Error; err != nil {
		t.Fatal(err)
	}

	post := map[string]string{
		"email": "test_user_1@user.com",
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	var result UserItemResultDto
	w := sendRequest(t, UriUser+UriUserGetByEmail, "POST", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, User.ID, result.ID)
}

func TestCreateUser_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)

	var result SuccessResponseDto
	post := map[string]string{
		"email":    "test_user_1@user.com",
		"password": "123123",
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	w := sendRequest(t, UriUser, "POST", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPutUserItem_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)

	now := time.Now()
	User := User{Email: "test_user_1@user.com", Password: "123123"}

	if err := db.Create(&User).Error; err != nil {
		t.Fatal(err)
	}

	User.DeletedAt = now
	jsonData, err := json.Marshal(User)
	if err != nil {
		panic(err)
	}
	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, User.ID.String()), "PUT", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	updatedUser, err := GetOneById(RequestUserIdDTO{
		ID: User.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, User.ID, updatedUser.ID)
}

func TestPatchUserItem_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)

	User := User{
		Email:    "test_user_1@user.com",
		Password: "123123",
	}

	if err := db.Create(&User).Error; err != nil {
		t.Fatal(err)
	}

	User.Email = "test_user_2@user.com"

	jsonData, err := json.Marshal(User)
	if err != nil {
		panic(err)
	}

	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, User.ID.String()), "PATCH", bytes.NewBuffer(jsonData), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	updatedUser, err := GetOneById(RequestUserIdDTO{
		ID: User.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, User.Email, updatedUser.Email)
}

func TestDeleteUserItem_SuccessfulResult(t *testing.T) {
	clearDbTableUser(t)

	User := User{Email: "test_user_1@user.com"}

	if err := db.Create(&User).Error; err != nil {
		t.Fatal(err)
	}

	var result SuccessResponseDto
	w := sendRequest(t, fmt.Sprintf(UriUser+UriUserGetByIdS, User.ID.String()), "DELETE", nil, &result)
	assert.Equal(t, http.StatusOK, w.Code)
	deletedUser, err := GetOneById(RequestUserIdDTO{
		ID: User.ID.String(),
	})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, uuid.Nil, deletedUser.ID)
}

// === Sys
func clearDbTableUser(t *testing.T) {
	if err := db.Exec("truncate table Users restart identity cascade").Error; err != nil {
		utils.Dump(err)
		t.Fatal(err)
	}
}

func sendRequest(
	t *testing.T,
	uri string,
	method string,
	body io.Reader,
	result any,
) *httptest.ResponseRecorder {

	//Init

	router := gin.Default()
	InitUserRoutes(router)

	// Creating test request
	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		t.Fatal(err)
	}

	//Sending test request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Parsing result

	err = json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	return w
}

func createUsers(length int) ([]User, error) {
	var Users []User
	for i := 1; i <= length; i++ {
		User := User{
			Email:    fmt.Sprintf("test_user_%d@user.com", i),
			Password: "123123",
		}
		if err := db.Create(&User).Error; err != nil {
			return Users, err
		}
		Users = append(Users, User)
	}
	return Users, nil
}
