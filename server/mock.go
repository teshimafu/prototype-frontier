package server

import (
	"errors"
	"strings"
)

//MockMiddleware no use firebase
type MockMiddleware struct{}

//InitFirebase is mock setup func
func (midd *MockMiddleware) InitFirebase() error {
	return nil
}

//VerifyIDToken is mock verify func
func (midd *MockMiddleware) VerifyIDToken(idToken string) (*AuthToken, error) {
	if strings.Contains(idToken, "invalid") {
		return nil, errors.New("invalid tenant id")
	}
	token := &AuthToken{
		UID:    idToken,
		Claims: map[string]interface{}{"email": "dummy@gmail.com", "name": "dummy_name"},
	}
	return token, nil
}
