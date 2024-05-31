package asteroid

/*
	context		Manejo de calncelación y plazos de las operaciones.
	errors		Manejo de errores.
	bson		Manejar documentos BSON, formato que usa MongoDB para almacenar documentos.
	primitive	Manejar tipos BSON primitivos, como ObjectID.
*/

import (
	"errors"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Declaramos la variable collection de tipo mongo.Collection, una estructura en el paquete "go.mongodb.org/mongo-driver/bson/mongo",
	que proporciona métodos para realizar operaciones CRUD en esa colección.
*/
type Repository struct {
	collection *mongo.Collection
}


/*
	Creamos un nuevo repositorio, recibe una base de datos y asigna asteroidsi
	Devolvemos un puntero al repositorio.
	Dentro de la funcion asignamos la colección asteroids pasado como parámetro a collection del repositorio.

*/
func	NewRepository(db *mongo.Database) *Repository {
	return	&Repository{
		collection: db.Collection("asteroids"),
	}
}

/*
	CREATE
	Método perteneciente al struct Repository y por tanto con permiso de acceso a sus campos y métodos.
	Recibe un argumento, asteroid, y generamos un ID uúnico para el asteroide y lo convertimos en cadena.
	ObjectID es el tipo que usa MongoDB com identificad único, lo convertimo a hexadecimal para hacerlo legible

*/
func (r *Repository) Create(asteroid Asteroid) (interface{}, error) {
	asteroid.ID = primitive.NewObjectID()
	result, err := r.collection.InsertOne(context.TODO(), asteroid)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

/*
	GetAll

	Un cursor en MongoDB es un putero a los resultados de una consulta, perminte iterar sobre ellos de forma eficiente.
	Cargamos todos los elementos en el slice asteroids

*/
func (r *Repository) GetAll() ([]Asteroid, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var asteroids []Asteroid
	if err = cursor.All(context.TODO(), &asteroids); err != nil {
		return nil, err
	}
	return asteroids, nil
}

/*
	GetByID

*/
func	(r *Repository) GetByID(id string) (Asteroid, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Asteroid{}, err
	}
	var asteroid Asteroid
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&asteroid)
	if err != nil {
		return Asteroid{}, errors.New("asteroid not foud")
	}
	return asteroid, nil
}

/*
	UPDATE
*/
func	(r *Repository) Update(id string, asteroid Asteroid) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": asteroid,
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": oid}, update)
	return err
}

/*
	DELETE
	Convertimos el id a ObjectID, tipo utilizado por MongoDB, para sus id únicos
*/
func	(r *Repository) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if result.DeletedCount == 0 {
		return errors.New("asteroid not found")
	}
	return nil
}

