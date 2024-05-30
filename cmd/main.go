package main

/*
	log para registrar  mensajes en la consola
	net/http para manejar solicictudes y respuestas HTTP.
*/
import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ALT-go-challenge/internal/asteroid"
)

/*
	Creamos una nueva instancia del repo, del handler(recibe el repo).
	Creamos un nuevo enrutador en r, con la función NewRouter del paquete Gorila Mux.
		- El enrutador (r) es una estructura que gestiona cómo se manejan las solicitudes HTTP entrantes.
		- HandleFunc es un método enrutador que define una ruta y la asocia con un método del handler.
		- Methods especifica el método HTTP(POST, GET, PATCH, DELETE)
	Inicilizamos el sevidor, con htt.ListenAndServer( en el puerto 8080).
	log.Fatal registra cualquier error en el sever.
*/
func main() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	db:= client.Database("asteroidsdb")
	repo := asteroid.NewRepository(db)
	handler := asteroid.NewHandler(repo)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/asteroides", handler.CreateAsteroid).Methods("POST")
	r.HandleFunc("/api/v1/asteroides", handler.GetAllAsteroids).Methods("GET")
	r.HandleFunc("/api/v1/asteroides/{id}", handler.GetAsteroidByID).Methods("GET")
	r.HandleFunc("/api/v1/asteroides/{id}", handler.UpdateAsteroid).Methods("PATCH")
	r.HandleFunc("/api/v1/asteroides/{id}", handler.DeleteAsteroid).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
