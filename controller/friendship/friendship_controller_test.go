package friendship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"friends_management_v2/services/friendship"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMakeFriendController(t *testing.T) {
	// Given
	testCase := []struct {
		name      string
		input     RequestFriend
		errorBody string
	}{
		{
			name: "Make Friend Success",
			input: RequestFriend{
				Friends: []string{
					"tony1@gmail.com",
					"tony2@gmail.com",
				},
			},
			errorBody: "",
		},
		{
			name: "Request Invalid",
			input: RequestFriend{
				Friends: []string{
					"quangbui1404@gmail.com",
				},
			},
			errorBody: "Request Invalid",
		},
		{
			name: "Email Invalid",
			input: RequestFriend{
				Friends: []string{
					"abc",
					"xyz",
				},
			},
			errorBody: "Email Invalid Format",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			values := map[string][]string{"friends": tc.input.Friends}

			frienshipMork := new(friendship.FrienshipMockService)
			if len(tc.input.Friends) == 2 {
				frienshipMork.On("MakeFriend", friendship.ServiceFrienshipInput{RequestEmail: tc.input.Friends[0], TargetEmail: tc.input.Friends[1]}).Return(nil)
			} else {
				frienshipMork.On("MakeFriend", friendship.ServiceFrienshipInput{RequestEmail: tc.input.Friends[0]}).Return(nil)
			}

			jsonValue, _ := json.Marshal(values)
			c.Request, _ = http.NewRequest("POST", "http://localhost:3000/add-friends", bytes.NewBuffer(jsonValue))
			c.Request.Header.Set("Content-Type", "application/json")
			// When

			MakeFriendController(c, frienshipMork)

			// Then
			var actualResult map[string]interface{}
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &actualResult)
			fmt.Println("Tony Quang ne", actualResult)
			if val1, ok1 := actualResult["success"]; ok1 {
				assert.Equal(t, val1, true)
			} else if val2, ok2 := actualResult["error"]; ok2 {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
				assert.Equal(t, val2, tc.errorBody)
			}

		})
	}

}
