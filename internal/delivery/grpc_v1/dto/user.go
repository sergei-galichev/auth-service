package dto

import (
	"regexp"
	"strings"
)

type UserCreateDTO struct {
	Email           string
	Password        string
	ConfirmPassword string
	Role            string
	AdminKey        string
}

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func IsValidEmail(email string) bool {
	return regexpEmail.MatchString(strings.ToLower(email))
}

func CheckPasswordsMatch(pass, confirmPass string) bool {
	return pass == confirmPass
}
