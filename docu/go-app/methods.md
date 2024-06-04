### Documentación de Ejemplos de Uso de la API

Esta documentación proporciona ejemplos claros y concisos de cómo utilizar los métodos disponibles en tu API. Incluye una explicación de las banderas utilizadas en los comandos `curl`.

### Ejemplos de Comandos `curl`

#### 1. Crear un Nuevo Asteroide (POST)

```sh
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "Asteroid X",
  "diameter": 1234.56,
  "discovery_date": "2024-05-01",
  "observations": "Second observation",
  "distances": [
    {
      "date": "2024-05-01",
      "distance": 400000
    }
  ]
}' http://localhost:8080/api/v1/asteroides
```

- **`-X POST`**: Especifica el método HTTP a utilizar (en este caso, POST).
- **`-H "Content-Type: application/json"`**: Establece el encabezado de la solicitud indicando que el cuerpo de la solicitud está en formato JSON.
- **`-d '{...}'`**: Especifica el cuerpo de la solicitud en formato JSON.
- **`http://localhost:8080/api/v1/asteroides`**: La URL del endpoint donde se envía la solicitud.

#### 2. Obtener Todos los Asteroides (GET)

```sh
curl http://localhost:8080/api/v1/asteroides
```

- **`http://localhost:8080/api/v1/asteroides`**: La URL del endpoint para obtener todos los asteroides. El método por defecto es GET.

#### 3. Obtener un Asteroide por ID (GET)

```sh
curl http://localhost:8080/api/v1/asteroides/{id}
```

- **`{id}`**: Reemplaza `{id}` con el ID real del asteroide que deseas obtener.
- **`http://localhost:8080/api/v1/asteroides/{id}`**: La URL del endpoint para obtener un asteroide específico por su ID.

#### 4. Actualizar un Asteroide por ID (PATCH)

```sh
curl -X PATCH -H "Content-Type: application/json" -d '{
  "name": "Updated Asteroid",
  "diameter": 5678.90
}' http://localhost:8080/api/v1/asteroides/{id}
```

- **`-X PATCH`**: Especifica el método HTTP a utilizar (en este caso, PATCH).
- **`-H "Content-Type: application/json"`**: Establece el encabezado de la solicitud indicando que el cuerpo de la solicitud está en formato JSON.
- **`-d '{...}'`**: Especifica el cuerpo de la solicitud en formato JSON.
- **`http://localhost:8080/api/v1/asteroides/{id}`**: La URL del endpoint para actualizar un asteroide específico por su ID.

#### 5. Eliminar un Asteroide por ID (DELETE)

```sh
curl -X DELETE http://localhost:8080/api/v1/asteroides/{id}
```

- **`-X DELETE`**: Especifica el método HTTP a utilizar (en este caso, DELETE).
- **`http://localhost:8080/api/v1/asteroides/{id}`**: La URL del endpoint para eliminar un asteroide específico por su ID.

