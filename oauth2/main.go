package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConfig *oauth2.Config

func init() {
	oauthConfig = &oauth2.Config{
		RedirectURL:  "localhost:8080/callback",
		ClientID:     "60848798350-v9dqmsun419pgk5ie3tj6hlk7hn8uqbu.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"http://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	http.HandleFunc("/login", handleGoogleLogin)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	randStateString := "asdasdlkajsdfkasjdfvn,mn@#@$"
	url := oauthConfig.AuthCodeURL(randStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
