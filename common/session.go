package common

import (
	"fmt"
	"strings"

	"kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

type AuthSession struct {
	repos.BaseEntity `json:",inline"`
	UserId           repos.ID `json:"user_id"`
	UserEmail        string   `json:"user_email"`
	UserVerified     bool     `json:"user_verified"`
	LoginMethod      string   `json:"login_method"`
}

type Json map[string]any

func NewSession(
	UserId repos.ID,
	UserEmail string,
	UserVerified bool,
	LoginMethod string,
) *AuthSession {
	id, e := functions.CleanerNanoid(28)
	if e != nil {
		panic(fmt.Errorf("could not get cleanerNanoid()"))
	}
	sessionId := fmt.Sprintf("ses-%s", strings.ToLower(id))
	s := &AuthSession{
		UserId:       UserId,
		UserEmail:    UserEmail,
		UserVerified: UserVerified,
		LoginMethod:  LoginMethod,
	}
	s.SetId(repos.ID(sessionId))
	return s
}