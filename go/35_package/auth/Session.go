package auth

func SetSession(token string) string {

	if token != "" {
		return ""
	} else {
		return "Session is set"
	}
}

func GetSession() string {
	return "Logged in"
}
