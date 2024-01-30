package repo

import "sophliteos/global"

func QueryUserWithName(username string) (*User, error) {
	var user User

	db := global.DB.Model(&User{}).Where("user_name = ?", &username).First(&user)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func QueryUserWithToken(token string) (*User, error) {
	var user User
	db := global.DB.Model(&User{}).Where("token = ?", &token).First(&user)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user *User) {
	values := map[string]interface{}{
		"token":       user.Token,
		"expire_time": user.ExpireTime,
	}
	global.DB.Model(&User{}).Where("id = ?", user.ID).Update(values)
}

func SaveUser(user *User) {
	global.DB.Model(&User{}).Save(user)
}
