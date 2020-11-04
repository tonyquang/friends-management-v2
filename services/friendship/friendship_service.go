package friendship

import (
	"errors"
	"fmt"
	"friends_management_v2/services/user"

	"gorm.io/gorm"
)

type FrienshipServices interface {
	AddFriends(input ServiceFrienshipInput) error
	GetFriendListOfAnUser(user user.Users) ([]string, error)
	GetMutualFriendsList(input ServiceFrienshipInput) ([]string, error)
	SubscribeUpdate(input ServiceFrienshipInput) error
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
	friendship, err := CheckFriendship(m.dbconn, firstUser, secondUser)

	if err != nil {
		return err
	}

	if friendship != nil {
		if friendship.IsFriend == true {
			return errors.New("Friendship was exist")
		}

		if friendship.UpdateStatus < 0 {
			return errors.New("Blocked Add Friend")
		}

		rs := m.dbconn.Model(&Friendship{}).Where("first_user IN ? AND second_user IN ?", []string{firstUser, secondUser}, []string{firstUser, secondUser}).Update("is_friend", true)
		if rs.Error != nil {
			return err
		}
		return nil
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

	rs := m.dbconn.Create(&Friendship{FirstUser: firstUser, SecondUser: secondUser, IsFriend: true, UpdateStatus: 0})

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

//SubscribeUpdate Update Subscribe
func (m *FriendshipManager) SubscribeUpdate(input ServiceFrienshipInput) error {
	listUsers := []string{input.First_user, input.Second_user}

	IsExist, err := user.CheckUserExist(m.dbconn, listUsers)

	if err != nil {
		return err
	}

	if IsExist == false {
		return errors.New("User Not Exist")
	}

	friendship, err := CheckFriendship(m.dbconn, input.First_user, input.Second_user)
	if err != nil {
		return err
	}

	if friendship != nil {
		stm := `UPDATE friendships as f0
		SET update_status = 
		(select 
			case 
				when f.update_status = 0 then 
					case 
						when f.first_user = @fuser then 1 else 2 
					end 
				else 
					case
						when f.first_user = @fuser then
							case
								when f.update_status <> 1 then 3 else 1
							end
						else
							case
								when f.update_status <> 2 then 3 else 2
							end
					end
			END AS update_status_code 
		from friendships as f where f.first_user IN (@fuser,@suser) AND f.second_user IN (@fuser,@suser)) 
		WHERE f0.first_user IN (@fuser,@suser) AND f0.second_user IN (@fuser,@suser)`

		rs := m.dbconn.Exec(stm, map[string]interface{}{"fuser": input.First_user, "suser": input.Second_user})

		if rs.Error != nil {
			return err
		}
	} else {
		rs1 := m.dbconn.Create(&Friendship{FirstUser: input.First_user, SecondUser: input.Second_user, IsFriend: false, UpdateStatus: 1})
		if rs1.Error != nil {
			return err
		}
	}

	return nil
}

//Check Connection Between Two User
func CheckFriendship(dbconn *gorm.DB, firstUser, secondUser string) (*Friendship, error) {
	friendship := Friendship{}
	rs := dbconn.Where("first_user IN ? AND second_user IN ?", []string{firstUser, secondUser}, []string{firstUser, secondUser}).Find(&Friendship{}).Scan(&friendship)
	if rs.Error != nil {
		return nil, rs.Error
	}
	fmt.Println(friendship)
	return &friendship, nil
}
