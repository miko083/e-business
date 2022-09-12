package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"

	"consoleshop/database"
	m "consoleshop/models"

	"time"

	"github.com/golang-jwt/jwt"
)

const BackEndToRedirect = "http://localhost:8000"
const frontEndLink = "http://localhost:3000/?login_token="

func databaseSignupOrLogin(email string, firstName string, lastName string, accessToken string) string {
	var userToCheck m.User
	query := database.DBconnection.Where("email = ?", email).Limit(1).Find(&userToCheck)
	userExists := query.RowsAffected > 0

	type jwtUser struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	claims := &jwtUser{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(generateState()))
	if err != nil {
		log.Println("Error when creating token!")
	}

	if userExists {
		log.Println(email + "is already in DB.")
		userToCheck.AccessToken = accessToken
		userToCheck.LoginToken = t
		userToCheck.IsLoggedIn = true
		database.DBconnection.Save(&userToCheck)
	} else {
		user := m.User{FirstName: firstName, LastName: lastName, Email: email, AccessToken: accessToken, LoginToken: t, IsLoggedIn: true}
		database.DBconnection.Create(&user)
	}

	url := frontEndLink + t
	return url
}

func generateState() string {
	// Random state - for safety
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

func JsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
