package check

import "regexp"

func CheckIsValidUsername(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9\.\-]{5,25}$`)
	return reg.MatchString(email)
}
