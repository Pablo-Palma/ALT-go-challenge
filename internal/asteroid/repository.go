package asteroid

/*
	Importamos paquetes estandard de go para el manejo de errores y la gesrion de Mutex
	(Mutual Exclusion), para evitar condiciones de carrera.
*/

import (
	"errors"
	"sync"
)

/*
	Declaramos una variable de tipo sync.Mutex del paquete sync para sincronizar acceso a los datos.
	y asteroids, una variable de tipo mapa, donde string(ID del asteroide) es la *KEY* y el *VALUE*
	es de tipo Asteroide(estructura definida en /internal/model.go)
*/
type Repository struct {
	mu			sync.Mutex
	asteroids	map[string]Asteroid
}


/*
	La función crea y devuelve un puntero a Repository (una estructura de tipo Repository)
	y además inicializa el campo asteroids con un mapa vacío.
	make: se usa para crear mapas, slices y canales.acepta tres argumentos:
		- tipo de estructura (slice, mapa, canal).
		- Longitud inicial (opcional para mapas y canales).
		- Capacidad (opcional para slices).
*/
func	NewRepository() *Repository {
	return	&Repository{
		asteroids: make(map[string]Asteroid),
	}
}

/*
	CREATE

	Create es un método del tipo Repositroy (r *Repository), r es el receptor del método y un puntero a Repository.
	En GO, un método de tipo "X", puede acceder y modificar variables de "X", siendo "X", una estructura.
	defer: Retrasa la ejecución del Unlock(), hasta  que la función termine.

	
*/
func (r *Repository) Create(asteroid Asteroid) {
	r.mu.Lock()									//Bloquear el mutex para acceso seguro.
	defer	r.mu.Unlock()						//Desbloquear al final de la función.
	r.asteroids[asteroid.ID] = asteroid			//Añade un asteroide al mapa, usando su `ID` como clave.
}

/*
	GetAll

	Devuelve un slice de Asteroid.
		Slice: Secuencia dinámica de elementos del mismo tipo, Internamente es una referencia a un array.
	En Go el for con range, puede devolver 2 valores: índice/valor, con "_, " ignoramos el índice.
*/
func (r *Repository) GetAll() []Asteroid {
	r.mu.Lock()
	defer	r.mu.Unlock()
	asteroids := make([]Asteroid, 0, len(r.asteroids))	// Creamos un slice de asteroides, con capacidad igual
														// a la longitud del mapa asteroids.
	for _, asteroid := range r.asteroids {				// Iteramos sobre cada valor("_, "), en el mapa asteroids.
		asteroids = append(asteroids, asteroid)			// Recorremos el mapa añadiendo cada asteroide al slice
	}
	return asteroids
}

/*
	GetByID

	Toma un string que representa el id, devuleve Asteroid si lo encuentra o un error.
*/
func	(r *Repository) GetByID(id string) (Asteroid, error) {
	r.mu.Lock()
	defer	r.mu.Unlock()
	asteroid, exists := r.asteroids[id]							// Buscamos el asteroide por su id.->True or false.
	if !exists {
		return Asteroid{}, errors.New("asteroid not found")		// Devolvemos el asteroide vacío, y un error.
	}
	return asteroid, nil										// Devolvermos el asteroide, y nil como error
																// para indicar operacion exitosa.
}

/*
	UPDATE
*/
func	(r *Repository) Update(id string, asteroid Asteroid) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.asteroids[id]
	if !exists {
		return errors.New("asteroid not found")
	}
	r.asteroids[id] = asteroid
	return nil
}

/*
	DELETE
*/
func	(r *Repository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.asteroids[id]
	if !exists {
		return errors.New("asteroid not found")
	}
	delete(r.asteroids, id)
	return nil
}

