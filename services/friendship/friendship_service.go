package friendship

import (
	"errors"
	"friends_management_v2/services/user"

	"gorm.io/gorm"
)

type FrienshipServices interface {
	AddFriends(input ServiceFrienshipInput) error
	GetFriendListOfAnUser(user user.Users) ([]string, error)
	GetMutualFriendsList(input ServiceFrienshipInput) ([]string, error)
}

// FriendshipManager is the implementation of recurring service
type FriendshipManager struct {
	dbconn *gorm.DB
}

// NewFriendshipManager initializes recurring service
func NewFriendshipManager(dbconn *gorm.DB) *FriendshipManager {
	return &FriendshipManager{
		dbconn: dbconn,
	}
}

func (m *FriendshipManager) AddFriends(input ServiceFrienshipInput) error {
	firstUser := input.First_user
	secondUser := input.Second_user

	//Check Friends Connection Exist
	IsFriendship, err := CheckFriendship(m.dbconn, firstUser, secondUser)

	if err != nil {
		return err
	}

	if IsFriendship == true {
		return errors.New("Friendship was exist")
	}

	//Check user exits
	listUsers := []string{firstUser, secondUser}

	ok, err := user.CheckUserExist(m.dbconn, listUsers)

	if err != nil {
		return err
	}

	if ok == false {
		return errors.New("User Not Exist")
	}

	for i := 0; i < 2; i++ {

	}

	rs := m.dbconn.Create(&Friendship{FirstUser: firstUser, SecondUser: secondUser, IsFriend: true, UpdateStatus: false})

	if rs.Error != nil {
		return rs.Error
	}

	return nil
}

//GetFriendListOfAnUser
func (m *FriendshipManager) GetFriendListOfAnUser(ur user.Users) ([]string, error) {

	IsExist, err := user.CheckUserExist(m.dbconn, []string{ur.Email})

	if err != nil {
		return nil, err
	}

	if IsExist == false {
		return nil, errors.New("User Not Exist")
	}

	stm := `SELECT f1.second_user friend FROM friendships as f1 WHERE f1.first_user = ? UNION SELECT f2.first_user friend FROM friendships as f2 WHERE f2.second_user = ?`

	listFriend := []string{}

	rs := m.dbconn.Raw(stm, ur.Email, ur.Email).Scan(&listFriend)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return listFriend, nil
}

//GetMutualFriendsList
func (m *FriendshipManager) GetMutualFriendsList(input ServiceFrienshipInput) ([]string, error) {

	listUsers := []string{input.First_user, input.Second_user}

	IsExist, err := user.CheckUserExist(m.dbconn, listUsers)

	if err != nil {
		return nil, err
	}

	if IsExist == false {
		return nil, errors.New("User Not Exist")
	}

	stm := `SELECT UserAFriends.friend FROM
	(
	 SELECT f1.second_user friend FROM friendships as f1 WHERE f1.first_user = ?
		UNION 
	 SELECT f2.first_user friend FROM friendships as f2 WHERE f2.second_user = ?
	) AS UserAFriends
	JOIN  
	(
	  SELECT f1.second_user friend FROM friendships as f1 WHERE f1.first_user = ?
		UNION 
	  SELECT f2.first_user friend FROM friendships as f2 WHERE f2.second_user = ?
	) AS UserBFriends 
	ON  UserAFriends.friend = UserBFriends.friend`

	listMutualFriends := []string{}
	rs := m.dbconn.Raw(stm, input.First_user, input.First_user, input.Second_user, input.Second_user).Scan(&listMutualFriends)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return listMutualFriends, nil //ffdf
}

//Check Connection Between Two User
func CheckFriendship(dbconn *gorm.DB, firstUser, secondUser string) (bool, error) {
	rs := dbconn.Where("first_user IN ? AND second_user IN ?", []string{firstUser, secondUser}, []string{firstUser, secondUser}).Find(&Friendship{})
	if rs.Error != nil {
		return false, rs.Error
	}
	if rs.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
