package main

import (
	"fmt"

	auth "github.com/manish-npx/Go-Learning/auth"
)

func main() {
	fmt.Println("Package FOlder")
	Email := "manish@email.com"
	Password := "1234@Password"
	authenticating := auth.LoginWithCreditional(Email, Password)
	checkedTOken := auth.CheckLogin("token set")
	fmt.Println(authenticating, checkedTOken)

}
