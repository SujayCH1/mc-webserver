package validation

import "regexp"

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)

func ValidUsername(username string) bool {
	return usernameRegex.MatchString(username)
}
