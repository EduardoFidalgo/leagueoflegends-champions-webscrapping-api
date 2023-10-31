package routes

import (
	"goWebScrapping/controllers"
	"goWebScrapping/models"
	"log"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest(characterList []models.Character) {
	r := mux.NewRouter()
	r.HandleFunc("/champions", func(w http.ResponseWriter, r *http.Request) {
		controllers.Home(w, r, characterList)
	})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
