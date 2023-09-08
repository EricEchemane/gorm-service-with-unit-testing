package session

var sessionStore = map[string]string{}

func Get(session_id string) string {
	return sessionStore[session_id]
}

func Set(session_id string, username string) {
	sessionStore[session_id] = username
}
