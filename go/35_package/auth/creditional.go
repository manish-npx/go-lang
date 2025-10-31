package auth

func CheckLogin(token string) bool {
	if token != "" {
		return false
	} else {
		return true
	}
}

func LoginWithCreditional(email, password string) string {
	if email == "" || password == "" {
		return "Email and Password are required"
	}
	return "Successfully Login"

}
