package friendship

import (
	"errors"

	"friends_management_v2/services/user"

	"gorm.io/gorm"
)

type FrienshipServices interface {
	MakeFriend(input FrienshipServiceInput) error
	GetFriendsList(user user.Users) ([]string, error)
	GetMutualFriendsList(input FrienshipServiceInput) ([]string, error)
	Subscribe(input FrienshipServiceInput) error
	Block(input FrienshipServiceInput) error
	GetUsersReceiveUpdate(sender string, mentionedUsers []string) ([]string, error)
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

func (m *FriendshipManager) MakeFriend(input FrienshipServiceInput) error {
	requestor := input.RequestEmail
	target := input.TargetEmail

	// Check Friends Connection Exist
	friendship, err := m.checkFriendship(requestor, target)

	if err != nil {
		return err
	}

	if friendship != nil {
		if friendship.IsFriend == true {
			return errors.New("Friendship was exist")
		}

		if friendship.UpdateStatus != 3 {
			return errors.New("Blocked Add Friend")
		}
		return m.execUpdateMakeFriend([]string{requestor, target})
	}

	// Check user exits
	listUsers := []string{requestor, target}

	ok, err := m.checkUserExist(listUsers)

	if err != nil {
		return err
	}

	if ok == false {
		return errors.New("User Not Exist")
	}

	// When Make Friend First User will subscribe update to Second User
	return m.execCreateFriendConnection(requestor, target, true, 1)
}

// execMakeFriend execute query make friend
func (m *FriendshipManager) execCreateFriendConnection(requestor string, target string, isFr bool, updateStatus int) error {
	rs := m.dbconn.Create(&Friendship{FirstUser: requestor, SecondUser: target, IsFriend: isFr, UpdateStatus: updateStatus})
	return rs.Error
}

// execMakeFriend execute query update make friend if requestor subscribe target but not be friend together
func (m *FriendshipManager) execUpdateMakeFriend(input []string) error {
	rs := m.dbconn.Model(&Friendship{}).Where("first_user IN ? AND second_user IN ?", input, input).Update("is_friend", true)
	return rs.Error
}

// GetUserFriendList
func (m *FriendshipManager) GetFriendsList(ur user.Users) ([]string, error) {

	IsExist, err := m.checkUserExist([]string{ur.Email})

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

// GetMutualFriendsList
func (m *FriendshipManager) GetMutualFriendsList(input FrienshipServiceInput) ([]string, error) {

	listUsers := []string{input.RequestEmail, input.TargetEmail}

	IsExist, err := m.checkUserExist(listUsers)

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
	rs := m.dbconn.Raw(stm, input.RequestEmail, input.RequestEmail, input.TargetEmail, input.TargetEmail).Scan(&listMutualFriends)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return listMutualFriends, nil
}

// Subscribe Update Subscribe
func (m *FriendshipManager) Subscribe(input FrienshipServiceInput) error {
	listUsers := []string{input.RequestEmail, input.TargetEmail}

	IsExist, err := m.checkUserExist(listUsers)

	if err != nil {
		return err
	}

	if IsExist == false {
		return errors.New("User Not Exist")
	}

	friendship, err := m.checkFriendship(input.RequestEmail, input.TargetEmail)
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
				when f.update_status = 3 then 3
				else
					case
						when f.first_user = @fuser then
							case
								when f.update_status = 1 then 1 else 3
							end
						else
							case
								when f.update_status = 1 then 3 else 2
							end
					end
			END AS update_status_code 
		from friendships as f where f.first_user IN (@fuser,@suser) AND f.second_user IN (@fuser,@suser)) 
		WHERE f0.first_user IN (@fuser,@suser) AND f0.second_user IN (@fuser,@suser)`

		rs := m.dbconn.Exec(stm, map[string]interface{}{"fuser": input.RequestEmail, "suser": input.TargetEmail})

		if rs.Error != nil {
			return err
		}
	} else {
		rs1 := m.dbconn.Create(&Friendship{FirstUser: input.RequestEmail, SecondUser: input.TargetEmail, IsFriend: false, UpdateStatus: 1})
		if rs1.Error != nil {
			return err
		}
	}
	return nil
}

func (m *FriendshipManager) Block(input FrienshipServiceInput) error {
	listUsers := []string{input.RequestEmail, input.TargetEmail}

	IsExist, err := m.checkUserExist(listUsers)

	if err != nil {
		return err
	}

	if IsExist == false {
		return errors.New("User Not Exist")
	}

	friendship, err := m.checkFriendship(input.RequestEmail, input.TargetEmail)
	if err != nil {
		return err
	}
	if friendship != nil {
		stm := `UPDATE friendships as f0
		SET update_status = 
		(select 
			case 
				when f.update_status = 3 then
					case
						when f.first_user = @fuser then 2 else 1	
					end
				else
					case
						when f.update_status <> 0 then
							case
								when f.first_user = @fuser then
									case
										when f.update_status = 1 then 0 else 2
									end
								else
									case
										when f.update_status = 2 then 0 else 1
									end
							end
						else 0	
					end	
			END AS update_status_code 
		from friendships as f where f.first_user IN (@fuser,@suser) AND f.second_user IN (@fuser,@suser)) 
		WHERE f0.first_user IN (@fuser,@suser) AND f0.second_user IN (@fuser,@suser)`

		rs := m.dbconn.Exec(stm, map[string]interface{}{"fuser": input.RequestEmail, "suser": input.TargetEmail})

		if rs.Error != nil {
			return err
		}
	} else {

		rs1 := m.dbconn.Create(&Friendship{FirstUser: input.RequestEmail, SecondUser: input.TargetEmail, IsFriend: false, UpdateStatus: -1})
		if rs1.Error != nil {
			return err
		}
	}
	return nil
}

func (m *FriendshipManager) GetUsersReceiveUpdate(sender string, metion []string) ([]string, error) {
	listUsers := []string{sender}

	IsExist, err := m.checkUserExist(listUsers)

	if err != nil {
		return nil, err
	}

	if IsExist == false {
		return nil, errors.New("User Not Exist")
	}

	stm := `select
				f1.second_user friend
			from
				friendships as f1
			where
				f1.first_user = ?
				and (f1.update_status <> 1 AND f1.update_status <> 0)  
				and (f1.is_friend = true or f1.update_status = 2 or f1.update_status = 3)
			union
			select
				f2.first_user friend
			from
				friendships as f2
			where
				f2.second_user = ?
				and (f2.update_status <> 2 OR f2.update_status <> 0)  
				and (f2.is_friend = true or f2.update_status = 1 or f2.update_status = 3)`

	listFriend := []string{}

	rs := m.dbconn.Raw(stm, sender, sender).Scan(&listFriend)

	if rs.Error != nil {
		return nil, err
	}

	mentionValid := []string{}

	rsCheckMentionValid := m.dbconn.Raw("select email from users where email IN ?", metion).Scan(&mentionValid)

	if rsCheckMentionValid.Error != nil {
		return nil, err
	}

	listFriend = append(listFriend, mentionValid...)

	return listFriend, nil
}

// Check Connection Between Two User
func (m *FriendshipManager) checkFriendship(firstUser, secondUser string) (*Friendship, error) {
	friendship := Friendship{}
	rs := m.dbconn.Where("first_user IN ? AND second_user IN ?", []string{firstUser, secondUser}, []string{firstUser, secondUser}).Find(&Friendship{}).Scan(&friendship)
	if rs.Error != nil {
		return nil, rs.Error
	}
	if rs.RowsAffected <= 0 {
		return nil, nil
	}
	return &friendship, nil
}

func (m *FriendshipManager) checkUserExist(listUsers []string) (bool, error) {
	ur := user.NewUserManager(m.dbconn)
	return ur.CheckUserExist(listUsers)
}
