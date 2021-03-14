package server

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func createFireStoreJSON() {
	err := godotenv.Load(fmt.Sprintf("./.env"))
	if err != nil {
		fmt.Println("can't read .env")
	}
	fp, err := os.Create("./firebaseCredentials.json")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer fp.Close()

	file := fmt.Sprintf(` {
    "type": "%s",
    "project_id": "%s",
    "private_key_id": "%s",
    "private_key": "%s",
    "client_email": "%s",
    "client_id": "%s",
    "auth_uri": "%s",
    "token_uri": "%s",
    "auth_provider_x509_cert_url": "%s",
    "client_x509_cert_url": "%s"
}`,
		os.Getenv("FIREBASE_Type"),
		os.Getenv("FIREBASE_ProjectID"),
		os.Getenv("FIREBASE_PrivateKeyID"),
		os.Getenv("FIREBASE_PrivateKey"),
		os.Getenv("FIREBASE_ClientEmail"),
		os.Getenv("FIREBASE_ClientID"),
		os.Getenv("FIREBASE_AuthURI"),
		os.Getenv("FIREBASE_TokenURI"),
		os.Getenv("FIREBASE_AuthProviderX509CertURL"),
		os.Getenv("FIREBASE_ClientX509CertURL"),
	)

	_, err = fp.Write(([]byte)(file))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
