package controllers

import (
	"consoleshop/database"
	m "consoleshop/models"
	"encoding/json"
	"os"
)

func checkIfAuthenticated(bodyBytes []byte) bool {
	body := make(map[string]interface{})
	json.Unmarshal(bodyBytes, &body)
	email := body["user_email"].(string)
	loginTokenFromFront := body["login_token"].(string)
	var userToCheck m.User
	database.DBconnection.Where("email = ?", email).Limit(1).Find(&userToCheck)
	return loginTokenFromFront == userToCheck.LoginToken
}

func checkIfAdmin(bodyBytes []byte) bool {
	body := make(map[string]interface{})
	json.Unmarshal(bodyBytes, &body)
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPasswordFromRequest, ok := body["admin_token"]; ok {
		return adminPassword == adminPasswordFromRequest
	}
	return false
}
