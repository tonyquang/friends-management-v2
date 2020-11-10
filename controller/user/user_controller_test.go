package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"friends_management_v2/services/user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUserController(t *testing.T) {
	// Given
	testCase := []struct {
		name              string
		inputRequest      RequestCreateUser
		expectedErrorBody string
	}{
		{
			name:              "Create New User Success",
			inputRequest:      RequestCreateUser{"abc@gmail.com"},
			expectedErrorBody: "",
		},
		{
			name:              "Create New User Fail",
			inputRequest:      RequestCreateUser{"abc@gmail.com"},
			expectedErrorBody: "Any error",
		},
		{
			name:              "Invalid User Email",
			inputRequest:      RequestCreateUser{"abc"},
			expectedErrorBody: "Invalid Email",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			userMock := new(user.UserMockService)
			if tc.name == "Create New User Fail" {
				userMock.On("CreateNewUser", user.Users{Email: tc.inputRequest.Email}).Return(errors.New("Any error"))
			} else {
				userMock.On("CreateNewUser", user.Users{Email: tc.inputRequest.Email}).Return(nil)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			value := map[string]string{"email": tc.inputRequest.Email}
			jsonValue, _ := json.Marshal(value)
			c.Request, _ = http.NewRequest("POST", "/create-user", bytes.NewBuffer(jsonValue))

			// When
			CreateNewUserController(c, userMock)

			//Then
			var actualResult map[string]interface{}
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &actualResult)

			if val1, oke1 := actualResult["success"]; oke1 {
				assert.Equal(t, val1, true)
			} else if val2, oke2 := actualResult["error"]; oke2 {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, val2, tc.expectedErrorBody)
			}
		})
	}

}

// func TestGetListUsersController(t *testing.T) {
// 	// Given

// 	testCase := []struct {
// 		name                string
// 		expectedErrorBody   string
// 		expectedSuccessBody ResponeListUser
// 	}{
// 		{
// 			name: "Get List User Success",
// 			expectedSuccessBody: ResponeListUser{
// 				Count: 3,
// 				ListUsers: []string{
// 					"abc@gmail.com",
// 					"xyz@gmail.com",
// 					"tonyquang@gmail.com",
// 				},
// 			},
// 		},
// 		{
// 			name:              "Get List User Fail",
// 			expectedErrorBody: "Any error",
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			userMock := new(user.UserMockService)
// 			if tc.name == "Get List User Success" {
// 				userMock.On("GetListUser").Return(tc.expectedSuccessBody.ListUsers, nil)
// 			} else {
// 				userMock.On("GetListUser").Return(nil, errors.New(tc.expectedErrorBody))
// 			}

// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)
// 			c.Request, _ = http.NewRequest("GET", "/list-users", nil)

// 			// When
// 			GetListUsersController(c, userMock)

// 			// Then
// 			var actualResult map[string]interface{}
// 			body, _ := ioutil.ReadAll(w.Result().Body)
// 			json.Unmarshal(body, &actualResult)
// 			fmt.Println(tc.name)
// 			fmt.Println(actualResult)
// 			if val1, oke1 := actualResult["error"]; oke1 {
// 				assert.Equal(t, val1, tc.expectedErrorBody)
// 			} else {
// 				_, oke2 := actualResult["list_users"]
// 				_, oke3 := actualResult["count"]
// 				assert.Equal(t, oke2, true)
// 				assert.Equal(t, oke3, true)
// 				assert.Equal(t, 200, w.Result().StatusCode)
// 			}
// 		})
// 	}
// }
