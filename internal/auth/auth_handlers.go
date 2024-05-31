package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var userRepo *UserRepository

func InitializeUserRepo(repo *UserRepository) {
	userRepo = repo
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	userID, err := userRepo.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = userID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	storedUser, err := userRepo.GetByUsername(user.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	if !CheckPasswordHash(user.Password, storedUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie {
		Name:		"token", 
		Value:		token,
		Expires:	time.Now().Add(24 * time.Hour),
	})
}

/*
	Este método: http.HandlerFunc, convierte una funcion con la firma recibida en un 
	http.HandlerFunc, que es un adaptador para convertif funciones en manejadores HTTP compatibles con la interfaz `http.Handler`
*/
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims, err := VerifyJWT(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	w.Write([]byte("Hello, " + username))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	err := userRepo.DeleteByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
