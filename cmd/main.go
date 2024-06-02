package main

/*
	logrus				para registrar  mensajes en la consola
	net/http 			para manejar solicictudes y respuestas HTTP.
	gorilla/handlers 	nos permite habilitar un middleware CORS(Cross Origin Resource Sharing)
	para conectar SwaggerUI a nuestra app, y así poder hacer peticiones.
*/
import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ALT-go-challenge/internal/asteroid"
	"ALT-go-challenge/internal/auth"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

/*
	Creamos una nueva instancia del repo, del handler(recibe el repo).
	Creamos un nuevo enrutador en r, con la función NewRouter del paquete Gorila Mux.
		- El enrutador (r) es una estructura que gestiona cómo se manejan las solicitudes HTTP entrantes.
		- HandleFunc es un método enrutador que define una ruta y la asocia con un método del handler.
		- Methods especifica el método HTTP(POST, GET, PATCH, DELETE)
	Inicilizamos el sevidor, con http.ListenAndServer( en el puerto 8080).
	log.Fatal registra cualquier error en el sever.
*/
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Starting the application")
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Connected to MongoDB!")

	db:= client.Database("asteroidsdb")
	repo := asteroid.NewRepository(db)
	handler := asteroid.NewHandler(repo)

	userRepo := auth.NewUserRepository(db)
	auth.InitializeUserRepo(userRepo)

	r := mux.NewRouter()

	r.Use(prometheusMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API  de Registro de Asteroides en funcionamiento"))
	})

	// Rutas públicas
	r.HandleFunc("/api/v1/asteroides", handler.GetAllAsteroids).Methods("GET")
	r.HandleFunc("/api/v1/asteroides/{id}", handler.GetAsteroidByID).Methods("GET")

	//Rutas protegidas
	r.Handle("/api/v1/asteroides", auth.AuthMiddleware(http.HandlerFunc(handler.CreateAsteroid))).Methods("POST")
	r.Handle("/api/v1/asteroides/{id}", auth.AuthMiddleware(http.HandlerFunc(handler.UpdateAsteroid))).Methods("PATCH")
	r.Handle("/api/v1/asteroides/{id}", auth.AuthMiddleware(http.HandlerFunc(handler.DeleteAsteroid))).Methods("DELETE")

	//Rutas de autenticación
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.Handle("/protected", auth.AuthMiddleware(http.HandlerFunc(auth.ProtectedEndpoint))).Methods("GET")
	r.Handle("/deleteuser", auth.AuthMiddleware(http.HandlerFunc(auth.DeleteUser))).Methods("DELETE")

	r.Handle("/metrics", promhttp.Handler())

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	logrus.Info("Server running on port 8080")
	logrus.Info("Access the application at: http://localhost:8080/")
	logrus.Info("Access Prometheus at: http://localhost:9090/")
	logrus.Info("Access Grafana at: http://localhost:3000/")
	logrus.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
		next.ServeHTTP(w, r)
	})
}
