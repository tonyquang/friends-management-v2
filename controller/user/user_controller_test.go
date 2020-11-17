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
		scenario          string
		inputRequest      *RequestCreateUser
		expectedErrorBody string
	}{
		{
			scenario:          "Create New User Success",
			inputRequest:      &RequestCreateUser{"abc@gmail.com"},
			expectedErrorBody: "",
		},
		{
			scenario:          "Create New User Fail",
			inputRequest:      &RequestCreateUser{"abc@gmail.com"},
			expectedErrorBody: "Any error",
		},
		{
			scenario:          "Invalid User Email",
			inputRequest:      &RequestCreateUser{"abc"},
			expectedErrorBody: "Invalid Email",
		},
		{
			scenario:          "Invalid User Email",
			inputRequest:      &RequestCreateUser{"abc"},
			expectedErrorBody: "Invalid Email",
		},
		{
			scenario:          "Empty request body",
			expectedErrorBody: "BindJson Error, cause body request invalid",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			userMock := new(user.UserMockService)

			if tc.inputRequest != nil {
				if tc.scenario == "Create New User Fail" {
					userMock.On("CreateNewUser", user.Users{Email: tc.inputRequest.Email}).Return(errors.New("Any error"))
				} else {
					userMock.On("CreateNewUser", user.Users{Email: tc.inputRequest.Email}).Return(nil)
				}
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if tc.inputRequest != nil {
				value := map[string]string{"email": tc.inputRequest.Email}
				jsonValue, _ := json.Marshal(value)
				c.Request, _ = http.NewRequest("POST", "/create-user", bytes.NewBuffer(jsonValue))
			} else {
				c.Request, _ = http.NewRequest("POST", "/create-user", nil)
			}
			// When
			CreateNewUserController(c, userMock)

			//Then
			var actualResult map[string]interface{}
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &actualResult)

			if val1, oke1 := actualResult["success"]; oke1 {
				assert.Equal(t, 201, w.Result().StatusCode)
				assert.Equal(t, val1, true)
			} else if val2, oke2 := actualResult["error"]; oke2 {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, val2, tc.expectedErrorBody)
			}
		})
	}

}

func TestGetListUsersController(t *testing.T) {
	//Given
	testCase := []struct {
		scenario            string
		mockRespone         []string
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			scenario: "Get List User Success",
			mockRespone: []string{
				"1@gmail.com",
				"2@gmail.com",
				"3@gmail.com",
				"4@gmail.com",
			},
			expectedSuccessBody: `{"list_users":["1@gmail.com","2@gmail.com","3@gmail.com","4@gmail.com"],"count":4}`,
		},
		{
			scenario:          "Get List User Fail",
			mockError:         errors.New("Any error"),
			expectedErrorBody: `{"error":"Any error"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			mockUser := new(user.UserMockService)
			mockUser.On("GetListUser").Return(tc.mockRespone, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			GetListUsersController(c, mockUser)

			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult := string(body)

			if tc.scenario == "Get List User Success" {
				assert.Equal(t, 200, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}
}
