package authentication

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/labstack/echo/v4"
)

// URL for Google OAuth2
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// State used for AuthCodeURL
var stateGoogle = ""

// Google OAuth2 Config
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func AuthGoogleLogin(c echo.Context) error {

	stateGoogle = generateState()
	url := googleOauthConfig.AuthCodeURL(stateGoogle)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthGoogleCallback(c echo.Context) error {

	if c.FormValue("state") != stateGoogle {
		log.Println("Invalid OAuth Google state")
	}

	data, accessToken, err := getUserDataGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
	}

	if err != nil {
		log.Println("Invalid Data received: ", err.Error())
	}

	email := data["email"].(string)
	firstName := data["given_name"].(string)
	lastName := data["family_name"].(string)

	url := databaseSignupOrLogin(email, firstName, lastName, accessToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func getUserDataGoogle(code string) (map[string]interface{}, string, error) {

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", fmt.Errorf("Code exchange wrong: %s", err.Error())
	}

	// Prepare URL
	urlWithToken := oauthGoogleUrlAPI + token.AccessToken
	response, err := http.Get(urlWithToken)
	if err != nil {
		return nil, "", fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	// Read from response
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed read response: %s", err.Error())
	}

	// Convert to map
	var contentsAsJson = JsonToMap(string(contents))
	return contentsAsJson, token.AccessToken, nil
}
