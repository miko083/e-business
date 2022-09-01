package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/labstack/echo/v4"

	"consoleshop/database"
	m "consoleshop/models"

	"github.com/golang-jwt/jwt"
)

// URL for Google OAuth2
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

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

	// Random state
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Create and set cookie
	cookie := new(http.Cookie)
	cookie.Name = "oauthstate"
	cookie.Value = state
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)

	url := googleOauthConfig.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthGoogleCallback(c echo.Context) error {
	// Read oauthState from Cookie
	oauthState, err := c.Cookie("oauthstate")
	if err != nil {
		log.Printf("Failed: %s", err)
	}

	if c.FormValue("state") != oauthState.Value {
		log.Println("Invalid OAuth Google state")
	}

	data, err := getUserData(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
	}

	if err != nil {
		log.Println("Invalid Data received: ", err.Error())
	}

	email := data["email"].(string)
	var user m.User
	query := database.DBconnection.Where("email = ?", email).Limit(1).Find(&user)
	userExists := query.RowsAffected > 0

	if userExists {
		fmt.Println(email + "is already in DB.")
	} else {
		user = m.User{FirstName: data["given_name"].(string), LastName: data["family_name"].(string), Email: email}
		database.DBconnection.Create(&user)
	}

	type jwtUser struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	claims := &jwtUser{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("EBiznes123"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

func getUserData(code string) (map[string]interface{}, error) {

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	// Prepare URL
	urlWithToken := oauthGoogleUrlAPI + token.AccessToken
	response, err := http.Get(urlWithToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	// Read from response
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	// Convert to map
	contentsAsJson := jsonToMap(string(contents))
	return contentsAsJson, nil
}

func jsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
