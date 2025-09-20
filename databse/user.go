package databse

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var userList []User

func (u *User) Store() User {
	if u.ID != 0 {
		return *u
	}
	u.ID = len(userList) + 1
	userList = append(userList, *u)
	return *u

}

func FindUser(email, password string) *User {
	for _, u := range userList {
		if u.Email == email && u.Password == password {
			return &u
		}

	}
	return nil
}
