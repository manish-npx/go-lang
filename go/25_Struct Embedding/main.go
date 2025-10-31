package main

import "fmt"

// base type
type User struct {
	Name string
}

func (u User) Describe() string {
	return fmt.Sprintf("User: %s", u.Name)
}

// embedded type
type Admin struct {
	User
	Role string
}

func (a Admin) Describe() string {
	// override embedded method
	return fmt.Sprintf("Admin: %s (role: %s)", a.User.Name, a.Role)
}

// interface
type Describer interface {
	Describe() string
}

func main() {
	user := User{Name: "Manish"}
	admin := Admin{
		User: User{Name: "Manish"},
		Role: "System Administrator",
	}

	var d Describer

	d = user
	fmt.Println(d.Describe()) // User: Manish

	d = admin
	fmt.Println(d.Describe()) // Admin: Manish (role: System Administrator)
}
