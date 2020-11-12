package user

// func TestCreateNewUser(t *testing.T) {
// 	dbconn := utils.CreateConnection()
// 	testCase := []struct {
// 		name      string
// 		givenUser Users
// 		result    error
// 	}{
// 		{
// 			name:      "Create A New User Success",
// 			givenUser: Users{Email: "quang99@gmail.com"},
// 			result:    nil,
// 		},
// 		{
// 			name:      "User Is Already",
// 			givenUser: Users{Email: "quang99@gmail.com"},
// 			result:    errors.New("User is already!"),
// 		},
// 	}
// 	tx := dbconn.Begin()

// 	assert.NoError(t, utils.LoadFixture(tx, "../datatest/create_user.sql"))

// 	tx.SavePoint("sp2")
// 	for _, tt := range testCase {
// 		t.Run(tt.name, func(t *testing.T) {
// 			manager := NewUserManager(tx)

// 			rs := manager.CreateNewUser(tt.givenUser)

// 			assert.Equal(t, tt.result, rs)
// 		})
// 	}
// 	tx.RollbackTo("sp2")
// }

// func TestGetListUser(t *testing.T) {
// 	dbconn := utils.CreateConnection()
// 	manager := NewUserManager(dbconn)
// 	rs, err := manager.GetListUser()
// 	assert.Equal(t, tt.result, rs)
// }
