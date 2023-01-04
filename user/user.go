package user

import "github.com/an1l4/mocktest/iuser"

type User struct {
	IUser iuser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1,"sample test")
}
