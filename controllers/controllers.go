package controllers

import (
	"encoding/json"
	"goWebScrapping/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, characterList []models.Character) {
	json.NewEncoder(w).Encode(&characterList)
}
