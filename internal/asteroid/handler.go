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
)

type Handler struct {
	Repo *Repository
}

/*
	NewHandler, crea un Handler con el repsitorio dado.
	asignado el parámetro repo al campo Repo de la estructura Handler.
	En Go las llaves(`{}`) se utilizanpara incializar un struct con valores específicos.
*/
func	NewHandler(repo *Repository)	*Handler {
	return &Handler{Repo: repo}
}

/*
	CreateAsteroid
	- parámetros:
		w http.ResponseWriter:	interfaz para enviar respuesta http, contiene unicamente la respuesta..
		r *http.Request:		struct que representa una solicitud http, contiene toda la info de la solicitud.
			r:	es una variable de tipo htt.Request.
				htt.Request es una estructura de la librería net/http, que repesenta una solicitud HTTP.
			r.Body: es un campo de tipo io.ReadCloser, que contiene el cuerpod e la solicitud HTTP( ex: datos JSON
				enviados en una solicitud POST).
		En Go, los manejadores de soliccitudes deben tener esta firma (`func(w http.ResponseWriter, r *http.Request)`).
		para ser compatibles con el paquete `net/http`

	json.NewDecoder(r.Body).Decode(&asteroid):
		- json.NewDecoder(r.Body) Tomamos un io.Reader(r.body) y devolvemos un decodificador JSON que se puede leer desde io.Reader.
		- Decode(&asteroid): El método Decode del decodificador, intenta decodificar los datos JSON DEL io.Reader y almacenralos en la variable asteroid.

	- flujo:
		Decodificamos el cuerpo de la solicitud y lo almacenamos en la variable asteroid.
			A. Si desemboca en error devolvemos http 400 (Bad Request)
			B. Si funciona, llamamos al método create del repo para almacenar el asteroide,
			   establecemos el códiog de estado HTTP 201 (Created), en w.WriteHeader.
			   codificamos el asteroide en formato JSON y lo enviamos como respuesta.
	
*/
func	(h *Handler) CreateAsteroid(w http.ResponseWriter, r *http.Request) {
	var	asteroid Asteroid
	if err:= json.NewDecoder(r.Body).Decode(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.Repo.Create(asteroid)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asteroid)
}

/*
	GetAllAsteroids
	rellena la variable asteroides con todos los asteroides almacenados en el repo
	escribe el códiog de estado HTTP a 200 (OK)
	Codifica la lista de asteroides en formato JSON y la envía en la respuesta.
*/
func (h *Handler) GetAllAsteroids(w http.ResponseWriter, r *http.Request) {
	asteroids := h.Repo.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asteroids)
}

/*
	GetAsteroidByID
	Obtenemos el ID del asterodie de la ruta de solcitud utilizando `mux.Vars`
	LLamamos al método GetByID del repositorio para obtener el asteroide por su ID.
		A. Si el asteroide no se encuentra responemos con un error HTTP 404 (Not Found).
		B. Si se encuentra establecemos el código de estado HTTP a 200 (OK)
		   y codificamos el asteroide en formato JSON para enviarlo como respuesta.
	mux.Vars(r)["id"]:	mux es el paquete de enrutamiento de Gorilla
						Vars(r), es una función que toma una soliccitud HTTP('r')
						y devuleve un mapa (`map[string]string`) con los parámetro de la URL.
							ex: si la URL es `/api/v1/asteroides/123`, `mux.Vars(r)["id"]` devolverá "`123`",
							si tenemos esta ruta definida: /api/v1/asteroides/{id}, cuando llega la solcitud `/api/v1/asteroides/123`, 
							mux.Vars(r), devolverá un mapa {"id": "123"}, extraerá el valor "123" y lo asignará a la variable id.
*/
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

/*
	UpdateAsteroid
	Decodificamos el cuerpo de la función (JSON) y lo amacenamos en asteroid.
	llamamos al método Update del repositorio para actualizar el asteroide.
	Establecemos el código de estado HTTP a 200 (OK)
	Codificamos el asteroide actualizado en formato JSON y lo enviamos en la respuesta.

	Si la decodificación es exitosa Decode almacenrará los datos en asteroid y la funcion devolverá nil.
	
*/
func	(h *Handler) UpdateAsteroid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var	asteroid Asteroid
	if err := json.NewDecoder(r.Body).Decode(&asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err:= h.Repo.Update(id, asteroid); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asteroid)
}

/*
	DeleteAsteroid
	LLamamos al método Delete del repo para borrar el asteroide a partir de us id.
	Si funciona escribimos en el header el código HTTP 200 (OK)
*/
func	(h *Handler) DeleteAsteroid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Asteroid deleted successfully"}`))
}
