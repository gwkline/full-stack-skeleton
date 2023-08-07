package helpers

import (
	"github.com/badoux/checkmail"
)

func ValidateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	err = checkmail.ValidateHost(email)
	if err != nil {
		return err
	}

	return nil
}
