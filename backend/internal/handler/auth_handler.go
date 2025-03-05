package handler

import (
	"encoding/json"
	"net/http"

	"github.com/smokeyghost/QuickyPaste/internal/database"
	"github.com/smokeyghost/QuickyPaste/internal/models"
	"github.com/smokeyghost/QuickyPaste/internal/utils"
	"gorm.io/gorm"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	hashedPassword, err := utils.HashPassword(creds.Password)
	if err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "System Error. Failed to hash the password."))
		return
	}

	user := models.User{Username: creds.Username, Password: hashedPassword}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusConflict, "User already exist"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	var user models.User
	if err := database.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.WriteAPIError(w, utils.NewAPIError(http.StatusUnauthorized, "User not found."))
			return
		}

		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	if !utils.CheckPasswordHash(creds.Password, user.Password) {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusUnauthorized, "Invalid username or password."))
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Failed to generate token."))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
