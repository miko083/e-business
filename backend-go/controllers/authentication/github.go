package authentication

import (
	"context"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"github.com/labstack/echo/v4"
)

// URL for Github OAuth2
const oauthGithubUrlAPI = "https://api.github.com/user"

// State used for AuthCodeURL
var stateGithub = ""

// Github OAuth2 Config
var githubOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/github/callback",
	ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
	Scopes: []string{
		"read:user",
		"user:email",
	},
	Endpoint: github.Endpoint,
}

func AuthGithubLogin(c echo.Context) error {

	stateGithub = generateState()
	url := githubOauthConfig.AuthCodeURL(stateGithub)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthGithubCallback(c echo.Context) error {

	if c.FormValue("state") != stateGithub {
		log.Println("Invalid OAuth Gituhb state")
	}

	data, accessToken, err := getUserDataGithub(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
	}

	if err != nil {
		log.Println("Invalid Data received: ", err.Error())
	}

	email := data["email"].(string)
	firstName := ""
	lastName := data["login"].(string)

	url := databaseSignupOrLogin(email, firstName, lastName, accessToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func getUserDataGithub(code string) (map[string]interface{}, string, error) {

	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", fmt.Errorf("Code exchange wrong: %s", err.Error())
	}

	// Prepare URL
	requestGithub, err := http.NewRequest("GET", oauthGithubUrlAPI, nil)
	if err != nil {
		return nil, "", fmt.Errorf("Failed connecting to Github API: %s", err.Error())
	}
	requestGithub.Header.Add("Accept", "application/vnd.github.v3+json")
	requestGithub.Header.Add("Authorization", "token "+token.AccessToken)

	response, err := http.DefaultClient.Do(requestGithub)
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
	contentsAsJson := JsonToMap(string(contents))
	return contentsAsJson, token.AccessToken, nil
}
