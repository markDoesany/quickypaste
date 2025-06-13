package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smokeyghost/QuickyPaste/internal/database"
	"github.com/smokeyghost/QuickyPaste/internal/middleware"
	"github.com/smokeyghost/QuickyPaste/internal/models"
	"github.com/smokeyghost/QuickyPaste/internal/utils"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusBadRequest, "Cannot resolve request."))
		return
	}

	defer r.Body.Close()

	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusBadRequest, "Error in getting the user id"))
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusNotFound, "User not found"))
		return
	}

	note := models.Note{
		Content:  input.Content,
		Username: username,
		UserID:   user.ID,
	}

	result := database.DB.Create(&note)
	if result.Error != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Server error. Please try again."))
		return
	}

	note.ShareableLink = fmt.Sprintf("https://quickpaste.qappslock.com/notes/%s", note.ID)
	if err := database.DB.Save(&note).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Server error. Please try again."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Unable to fetch username"))
		return
	}

	var notes []models.Note
	if err := database.DB.Where("username = ?", username).Find(&notes).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Unable to fetch the note. Please try again."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var note models.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusNotFound, "The note was not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Server error. Unable to delete the note."))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Note deleted successfully"})
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Content string `json:"content"`
	}

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusBadRequest, "Unable to process the request."))
		return
	}

	defer r.Body.Close()

	var note models.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "The requested note does not exist"))
		return
	}

	note.Content = input.Content

	if err := database.DB.Save(&note).Error; err != nil {
		utils.WriteAPIError(w, utils.NewAPIError(http.StatusInternalServerError, "Unable to save changes."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
