package main

import (
	"github.com/getlantern/go-socks5"
)

type AuthStore struct {
	socks5.CredentialStore
	Username string
	Password string
}

func (s *AuthStore) Valid(user, password string) bool {
	//log.Errorf("AUTH: %s:%s", user, password)
	//fmt.Printf("AUTH: %s:%s", user, password)

	if user == s.Username && password == s.Password {
		return true
	}

	return false
}
