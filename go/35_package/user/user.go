package user

type User struct {
	Email    string
	Password string
}

func (u *User) UserDetails() {
	u.Email = "e@mail.com"
	u.Password = "password"
}
