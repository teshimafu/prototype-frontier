package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"

	"google.golang.org/api/option"
)

// MiddlewareFunc is middleware interface
type MiddlewareFunc interface {
	InitFirebase() error
	VerifyIDToken(idToken string) (*AuthToken, error)
}

// AuthToken is VerifyIDToken return
type AuthToken struct {
	UID    string                 `json:"uid,omitempty"`
	Claims map[string]interface{} `json:"-"`
}

// FirebaseMiddleware is firebase info
type FirebaseMiddleware struct {
	app *firebase.App
}

func authMiddleware(midd MiddlewareFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := midd.InitFirebase()
		if err != nil {
			log.Printf("error: %v\n", err)
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer ") {
			log.Printf("Invalid Request\n")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Request\n"})
			c.Abort()
			return
		}
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := midd.VerifyIDToken(idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "error verifying ID token\n"})
			c.Abort()
			return
		}
		log.Printf("uid: %v, email: %v, name: %v\n", token.UID, token.Claims["email"], token.Claims["name"])
		c.Set("AuthorizedUID", token.UID)
		c.Next()
	}
}

//InitFirebase is firebase set up
func (midd *FirebaseMiddleware) InitFirebase() error {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	var err error
	midd.app, err = firebase.NewApp(context.Background(), nil, opt)
	return err
}

//VerifyIDToken is return firebase setting app
func (midd *FirebaseMiddleware) VerifyIDToken(idToken string) (*AuthToken, error) {
	client, err := midd.app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	token, err := client.VerifyIDToken(context.Background(), idToken)
	authToken := &AuthToken{
		UID:    token.UID,
		Claims: token.Claims,
	}
	return authToken, err
}
