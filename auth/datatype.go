package auth

import (
	"shortener/common"
	"strings"
	"unicode"
)

type AuthorizationHeaderRequest struct {
	Authorization string `header:"authorization"`
}

func (req *AuthorizationHeaderRequest) Validate() *common.RestError {
	if !strings.HasPrefix(req.Authorization, "Bearer ") {
		return common.BadRequestError("authorization token should start with Bearer")
	}
	return nil
}

type RegisterUserRequest struct {
	LoginUserRequest
	Username string `json:"username" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req *LoginUserRequest) Validate() *common.RestError {
	err := common.BadRequestError("password should match minimum eight characters, at least one letter and one number")
	if len(req.Password) < 8 && len(req.Password) > 64 {
		return err
	}
	var haveLetter, haveDigit bool
	for _, let := range []rune(req.Password) {
		if unicode.IsDigit(let) {
			haveDigit = true

		}
		if unicode.IsLetter(let) {
			haveLetter = true
		}
	}
	if !haveLetter || !haveDigit {
		return err
	}
	return nil
}

type JWTClaims struct {
	Email string `json:"email"`
	Exp   uint64 `json:"exp"`
}
