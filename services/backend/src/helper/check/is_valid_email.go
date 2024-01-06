package check

import "net/mail"

func CheckIsValidEmail(email string) (bool, error) {
	_, addrErr := mail.ParseAddress(email)
	if addrErr != nil {
		return false, addrErr
	}

	// todo check if is valid email
	return true, nil
}
