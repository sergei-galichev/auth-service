package dto

import (
	"golang.org/x/crypto/bcrypt"
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

type UserLogoutDTO struct {
	AccessToken  string
	RefreshToken string
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

func ComparePassAndHash(pass string, passHash []byte) bool {
	return bcrypt.CompareHashAndPassword(passHash, []byte(pass)) == nil
}
