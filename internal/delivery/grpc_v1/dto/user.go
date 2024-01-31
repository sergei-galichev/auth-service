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

type UserLoginDTO struct {
	Email    string
	Password string
}

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func IsValidEmail(email string) bool {
	return regexpEmail.MatchString(strings.ToLower(email))
}

func ArePasswordsMatch(pass, confirmPass string) bool {
	return pass == confirmPass
}

func IsAdminRole(role string) bool {
	return role == "ADMIN"
}

func IsAdminPasswordValid(adminKey, password string) bool {
	return adminKey == password
}
