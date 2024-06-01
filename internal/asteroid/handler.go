package	asteroid


/*
	"encoding/json" nos permite trabajar con datos json, funciones clave:
		- json.NewDecoder: Decodifica JSON desde un `io.Reader` ex(`http.Request.Body`)
		- json.NewEncoder: Codifica JSON para escribirlos en un  `io.Reader` ex(`http.RequestWriter`)

	"net/http": proporciona herramientas para construir servidores HTTP y manejar solicitudes y respuestas HTTP.
	"mux": Enrutador HTTP.
		-mux.NewRouter()	Crea Enrutador.
		-mux.Vars(r)		Obtiene las variables de ruta como parámetros en la URL
*/

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	Repo *Repository
}

func	NewHandler(repo *Repository)	*Handler {
	return &Handler{Repo: repo}
}

func	(h *Handler) CreateAsteroid(w http.ResponseWriter, r *http.Request) {
	var	asteroid Asteroid
	if err:= json.NewDecoder(r.Body).Decode(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err:= ValidateAsteroid(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	asteroid.ID = primitive.NewObjectID()
	insertID, err := h.Repo.Create(asteroid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	asteroid.ID = insertID.(primitive.ObjectID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asteroid)
}

func (h *Handler) GetAllAsteroids(w http.ResponseWriter, r *http.Request) {
	asteroids, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asteroids)
}

func (h *Handler) GetAsteroidByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	asteroid, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asteroid)
}

func	(h *Handler) UpdateAsteroid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var	asteroid Asteroid
	if err := json.NewDecoder(r.Body).Decode(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := ValidateAsteroid(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := h.Repo.GetByID(id); err != nil {
		http.Error(w, "Asteroid not found", http.StatusNotFound)
		return
	}
	if err:= h.Repo.Update(id, asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asteroid)
}

func	(h *Handler) DeleteAsteroid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if _, err := h.Repo.GetByID(id); err != nil {
		http.Error(w, "Asteroid not found", http.StatusNotFound)
		return
	}
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Asteroid deleted successfully"}`))
}
