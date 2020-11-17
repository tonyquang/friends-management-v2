package friendship

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"friends_management_v2/services/friendship"
	"friends_management_v2/services/user"
	"friends_management_v2/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMakeFriendController(t *testing.T) {
	// Given
	testCase := []struct {
		// scenario
		name              string
		input             RequestFriend
		expectedErrorBody string
	}{
		{
			name: "Make Friend Success",
			input: RequestFriend{
				Friends: []string{
					"tony1@gmail.com",
					"tony2@gmail.com",
				},
			},
			expectedErrorBody: "",
		},
		{
			name: "Make Friend Fail",
			input: RequestFriend{
				Friends: []string{
					"someone1@gmail.com",
					"someone2@gmail.com",
				},
			},
			expectedErrorBody: "Any Error",
		},
		{
			name: "Not enough parameters",
			input: RequestFriend{
				Friends: []string{
					"quangbui1404@gmail.com",
				},
			},
			expectedErrorBody: "Request Invalid",
		},
		{
			name: "Same user",
			input: RequestFriend{
				Friends: []string{
					"quangbui1404@gmail.com",
					"quangbui1404@gmail.com",
				},
			},
			expectedErrorBody: "Request Invalid",
		},
		{
			name: "Email Invalid",
			input: RequestFriend{
				Friends: []string{
					"abc",
					"xyz",
				},
			},
			expectedErrorBody: "Email Invalid Format",
		},
		{
			name:              "Empty request body",
			expectedErrorBody: "BindJson Error, cause body request invalid",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			frienshipMock := new(friendship.FrienshipMockService)
			if tc.input.Friends != nil {
				if tc.name != "Make Friend Fail" {
					if len(tc.input.Friends) == 2 {
						frienshipMock.On("MakeFriend", friendship.FrienshipServiceInput{RequestEmail: tc.input.Friends[0], TargetEmail: tc.input.Friends[1]}).Return(nil)
					} else {
						frienshipMock.On("MakeFriend", friendship.FrienshipServiceInput{RequestEmail: tc.input.Friends[0]}).Return(nil)
					}
				} else {
					frienshipMock.On("MakeFriend", friendship.FrienshipServiceInput{RequestEmail: tc.input.Friends[0], TargetEmail: tc.input.Friends[1]}).Return(errors.New("Any Error"))
				}
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			values := map[string][]string{"friends": tc.input.Friends}
			jsonValue, _ := json.Marshal(values)
			c.Request, _ = http.NewRequest("POST", "/add-friends", bytes.NewBuffer(jsonValue))
			c.Request.Header.Set("Content-Type", "application/json")
			// When

			MakeFriendController(c, frienshipMock)

			// Then
			var actualResult map[string]interface{}
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &actualResult)

			if val1, ok1 := actualResult["success"]; ok1 {
				assert.Equal(t, val1, true)
			} else if val2, ok2 := actualResult["error"]; ok2 {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, val2, tc.expectedErrorBody)
			}

		})
	}

}

func TestGetFriendsList(t *testing.T) {

	// Given

	testCase := []struct {
		name                string
		input               user.Users
		mockRespone         []string
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			name:                "Get List Friends Success",
			input:               user.Users{Email: "abc@gmail.com"},
			mockRespone:         []string{"quang1@gmail.com", "xyz@gmail.com", "ok@yahoo.com"},
			expectedSuccessBody: `{"success":true,"friends":["quang1@gmail.com","xyz@gmail.com","ok@yahoo.com"],"count":3}`,
		},
		{
			name:              "Get List Friends Fail",
			input:             user.Users{Email: "abcxxx@gmail.com"},
			mockError:         errors.New("Any error"),
			mockRespone:       nil,
			expectedErrorBody: `{"error":"Any error"}`,
		},
		{
			name:              "Invalid Email",
			input:             user.Users{Email: "abc"},
			mockRespone:       nil,
			expectedErrorBody: `{"error":"Email Invalid Format"}`,
		},
		{
			name:              "Empty request body",
			expectedErrorBody: `{"error":"BindJson Error, cause body request invalid"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockFriendship := new(friendship.FrienshipMockService)
			mockFriendship.On("GetFriendsList", tc.input).Return(tc.mockRespone, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			values := map[string]string{"Email": tc.input.Email}
			jsonValue, _ := json.Marshal(values)
			c.Request, _ = http.NewRequest("POST", "/get-list-friends", bytes.NewBuffer(jsonValue))

			// When
			GetFriendsListController(c, mockFriendship)

			// Then
			var actualResult string
			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult = string(body)

			if tc.name == "Get List Friends Success" {
				assert.Equal(t, 200, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}

}

func TestGetMutualFriendsController(t *testing.T) {
	testCase := []struct {
		name                string
		requestInput        RequestFriend
		mockRespone         []string
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			name: "Get Mutual Friends Success",
			requestInput: RequestFriend{
				Friends: []string{
					"requestor@gmail.com",
					"target@gmail.com",
				},
			},
			mockRespone: []string{
				"mutual1@gmail.com",
				"mutual2@gmail.com",
				"mutual3@gmail.com",
			},
			expectedSuccessBody: `{"success":true,"friends":["mutual1@gmail.com","mutual2@gmail.com","mutual3@gmail.com"],"count":3}`,
		},
		{
			name: "Get Mutual Friends Fail",
			requestInput: RequestFriend{
				Friends: []string{
					"requestor@gmail.com",
					"target@gmail.com",
				},
			},
			mockError:         errors.New("Any error"),
			expectedErrorBody: `{"error":"Any error"}`,
		},
		{
			name: "Invalid Email",
			requestInput: RequestFriend{
				Friends: []string{
					"requestor",
					"target@gmail.com",
				},
			},
			expectedErrorBody: `{"error":"Email Invalid Format"}`,
		},
		{
			name: "requestor same target",
			requestInput: RequestFriend{
				Friends: []string{
					"target@gmail.com",
					"target@gmail.com",
				},
			},
			expectedErrorBody: `{"error":"Request Invalid"}`,
		},
		{
			name: "Not enough parameters",
			requestInput: RequestFriend{
				Friends: []string{
					"requestor@gmail.com",
				},
			},
			expectedErrorBody: `{"error":"Request Invalid"}`,
		},
		{
			name:              "Empty request body",
			expectedErrorBody: `{"error":"BindJson Error, cause body request invalid"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockFriendship := new(friendship.FrienshipMockService)
			if tc.requestInput.Friends != nil {
				if tc.name == "Not enough parameters" {
					mockFriendship.On("GetMutualFriendsList", friendship.FrienshipServiceInput{RequestEmail: tc.requestInput.Friends[0]}).Return(tc.mockRespone, tc.mockError)
				} else {
					mockFriendship.On("GetMutualFriendsList", friendship.FrienshipServiceInput{RequestEmail: tc.requestInput.Friends[0], TargetEmail: tc.requestInput.Friends[1]}).Return(tc.mockRespone, tc.mockError)
				}
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			jsonValue, _ := json.Marshal(tc.requestInput)

			c.Request, _ = http.NewRequest("POST", "/get-mutual-list-friends", bytes.NewBuffer(jsonValue))

			// When
			GetMutualFriendsController(c, mockFriendship)

			//Then

			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult := string(body)

			if tc.name == "Get Mutual Friends Success" {
				assert.Equal(t, 200, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}
}

func TestSubscribeController(t *testing.T) {

	// Given

	testCase := []struct {
		name                string
		inputRequest        RequestUpdate
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			name: "Subscribe Success",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
				Target:    "target@gmail.com",
			},
			expectedSuccessBody: `{"success":true}`,
		},
		{
			name: "Subscribe Fail",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
				Target:    "target@gmail.com",
			},
			mockError:         errors.New("Any error"),
			expectedErrorBody: `{"error":"Any error"}`,
		},
		{
			name: "Invalid Mail",
			inputRequest: RequestUpdate{
				Requestor: "requestor",
				Target:    "target",
			},
			expectedErrorBody: `{"error":"Email Invalid Format"}`,
		},
		{
			name: "Not enough parameters",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
			},
			expectedErrorBody: `{"error":"BindJson Error, cause body request invalid"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockFriendship := new(friendship.FrienshipMockService)
			mockFriendship.On("Subscribe", friendship.FrienshipServiceInput{RequestEmail: tc.inputRequest.Requestor, TargetEmail: tc.inputRequest.Target}).Return(tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonVal, _ := json.Marshal(tc.inputRequest)
			c.Request, _ = http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonVal))

			// When

			SubscribeController(c, mockFriendship)

			// Then

			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult := string(body)

			if tc.name == "Subscribe Success" {
				assert.Equal(t, 201, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}
}

func TestBlockController(t *testing.T) {
	// Given

	testCase := []struct {
		name                string
		inputRequest        RequestUpdate
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			name: "Block Success",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
				Target:    "target@gmail.com",
			},
			expectedSuccessBody: `{"success":true}`,
		},
		{
			name: "Block Fail",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
				Target:    "target@gmail.com",
			},
			mockError:         errors.New("Any error"),
			expectedErrorBody: `{"error":"Any error"}`,
		},
		{
			name: "Invalid Mail",
			inputRequest: RequestUpdate{
				Requestor: "requestor",
				Target:    "target",
			},
			expectedErrorBody: `{"error":"Email Invalid Format"}`,
		},
		{
			name: "Not enough parameters",
			inputRequest: RequestUpdate{
				Requestor: "requestor@gmail.com",
			},
			expectedErrorBody: `{"error":"BindJson Error, cause body request invalid"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockFriendship := new(friendship.FrienshipMockService)
			mockFriendship.On("Block", friendship.FrienshipServiceInput{RequestEmail: tc.inputRequest.Requestor, TargetEmail: tc.inputRequest.Target}).Return(tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonVal, _ := json.Marshal(tc.inputRequest)
			c.Request, _ = http.NewRequest("POST", "/block", bytes.NewBuffer(jsonVal))

			// When

			BlockController(c, mockFriendship)

			// Then

			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult := string(body)

			if tc.name == "Block Success" {
				assert.Equal(t, 201, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}
}

func TestGetUsersReceiveUpdateController(t *testing.T) {
	// Given
	testCase := []struct {
		name                string
		inputRequest        *RequestReceiveUpdate
		mockRespone         []string
		mockError           error
		expectedErrorBody   string
		expectedSuccessBody string
	}{
		{
			name: "Recvice Success",
			inputRequest: &RequestReceiveUpdate{
				Sender: "quang@gmail.com",
				Text:   "Hello world!, hi @buiminhquang@yahoo.com",
			},
			mockRespone: []string{
				"buiminhquang@yahoo.com",
				"tonyquangdeptrai@gmail.com",
				"tonyquang9x@gmail.com",
			},
			expectedSuccessBody: `{"success":true,"recipients":["buiminhquang@yahoo.com","tonyquangdeptrai@gmail.com","tonyquang9x@gmail.com"]}`,
		},
		{
			name: "Recvice Fail",
			inputRequest: &RequestReceiveUpdate{
				Sender: "quang@gmail.com",
				Text:   "Hello world!, hi @buiminhquang@yahoo.com",
			},
			mockError:         errors.New("Any error"),
			expectedErrorBody: `{"error":"Any error"}`,
		},
		{
			name: "Invalid Email",
			inputRequest: &RequestReceiveUpdate{
				Sender: "quang",
				Text:   "Hello world!, hi @buiminhquang@yahoo.com",
			},
			expectedErrorBody: `{"error":"Email Invalid Format"}`,
		},
		{
			name:              "Empty request body",
			expectedErrorBody: `{"error":"BindJson Error, cause body request invalid"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockFriendship := new(friendship.FrienshipMockService)
			if tc.inputRequest != nil {
				mentioned := utils.ExtractMentionEmail(tc.inputRequest.Text)
				mockFriendship.On("GetUsersReceiveUpdate", tc.inputRequest.Sender, mentioned).Return(tc.mockRespone, tc.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonVal, _ := json.Marshal(tc.inputRequest)
			c.Request, _ = http.NewRequest("POST", "/get-list-users-receive-update", bytes.NewBuffer(jsonVal))

			// When

			GetUsersReceiveUpdateController(c, mockFriendship)

			// Then

			body, _ := ioutil.ReadAll(w.Result().Body)
			actualResult := string(body)

			if tc.name == "Recvice Success" {
				assert.Equal(t, 200, w.Result().StatusCode)
				assert.Equal(t, tc.expectedSuccessBody, actualResult)
			} else {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, tc.expectedErrorBody, actualResult)
			}
		})
	}
}
