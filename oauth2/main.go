package main

import (
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
