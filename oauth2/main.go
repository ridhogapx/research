package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
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
		Scopes:       []string{"https://www.googleapis.com/auth/blogger"},
		Endpoint:     google.Endpoint,
	}
}

var randStateString = "asdasdlkajsdfkasjdfvn,mn@#@$"

func main() {
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleCallbackGoogle)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state")
	fmt.Println("Visit url:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallbackGoogle(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Content: %v", content)

}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != randStateString {
		return nil, errors.New("invalid oauth state" + state)
	}

	token, err := oauthConfig.Exchange(context.Background(), code)

	if err != nil {
		return nil, errors.New("code exchange failed")
	}

	oauthGoogleUrlAPI := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	client := oauthConfig.Client(context.Background(), token)

	response, err := client.Get(oauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		return nil, errors.New("failed getting user info")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.New("failed reading response body")
	}

	return contents, nil

}
