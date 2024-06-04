# GO APLICATION

En este proyecto inspeccionamos el mundo de las aplicaciones web usando el lenguaje de programación Go para desplegar una aplicación que nos permite desplegar una api contenerizada con docker, una aplicaición de código abierto que nos permite automatizar el despligue de aplicaciones dentro de contenedores de software.

Nuestra API almacena los datos tanto de los usuarios como de los objetos en una base de datos NoSQL, en este caso hemos elegido mongodb.

Nuestra api nos proporciona un sistema donde emplear metodos CRUD(Create Research Update Delete)
* Almacenar información sobre asteroides: incluyendo nombre, tamaño, composición, etc.
* Registrar la distancia del asteroide a la Tierra en diferentes fechas.
* Permitir la consulta de información sobre asteroides específicos.
* Permitir la actualización de la información de un asteroide.
* Eliminar información de asteroides que ya no son una amenaza.

Como funciones adicionales hemos implementado:

* Paginación y filtros.
* Sistema de autenticación y autorización para controlar el acceso a la API.
* Validación de entrada de datos para evitar errores y datos incorrectos.
* Monitoring: Uso de herramientas de trazas, logging y Monitorizacion de codigos de estado.
* Documentación: Hemos empleado OpenApi para describir, producir, consumir y visualizar servicios web RESTful.
* GitHub como repositorio de almacenamiento y control de versiones.
* Testing: Hemos incluido test unitarios, tanto en go como postman
* Frontend.

### Tecnologías Utilizadas

- **Lenguaje de Programación**: Go
- **Contenerización**: Docker
- **Base de Datos**: MongoDB (NoSQL)
- **Documentación de API**: OpenAPI
- **Control de Versiones**: GitHub
- **Testing**: Go, Postman
- **Monitoring y Logging**: Prometheus, Grafana
- **Frontend**: Astro.


## Estructura del Proyecto

```
.
├── cmd
│   └── main.go
├── docu
│   ├── *.md
├── internal
│   ├── asteroid
│   │   ├── handler.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── validate.go
│   └── auth
│       ├── auth_handlers.go
│       ├── jwt.go
│       ├── user.go
│       └── user_repository.go
├── prometheus
│   └── prometheus.yml
├── frontend
│   ├── README.md
│   ├── node_modules
│   ├── *.json
│   ├── public
│   ├── src
│   │   ├── components
│   │   │   ├── AsteroidBox.astro
│   │   │   ├── AsteroidList.astro
│   │   │   ├── Button.astro
│   │   │   └── Card.astro
│   │   ├── layouts
│   │   │   └── Layout.astro
│   │   ├── pages
│   │   │   ├── explore.astro
│   │   │   └── index.astro
│   │   └── env.d.ts
│   └── tsconfig.json
└── Dockerfile
└── go.mod
```

## Organización del README

<details>
<summary><strong>1. Descripción del Proyecto</strong></summary>

- Introducción
- Objetivo del Proyecto
- Características Principales

</details>

<details>
<summary><strong>2. Implementación de la Aplicación en Go</strong></summary>

#### Explicación:

- **cmd/main.go**: Contiene el punto de entrada de la aplicación.
- **internal/asteroid/**: Contiene el código relacionado con los asteroides.
  - **model.go**: Define la estructura (modelo) de los asteroides.
  - **repository.go**: Gestiona el almacenamiento y la recuperación de datos.
  - **handler.go**: Maneja las solicitudes HTTP relacionadas con los asteroides.
- **internal/router/**: Configura las rutas de la API.
- **Dockerfile**: Define cómo se construye la imagen Docker para la aplicación.
- **go.mod**: Archivo de configuración del módulo Go.

### 2. ¿Qué es un Módulo Go y Cómo Inicializarlo?

Un módulo Go es una colección de paquetes Go que se gestionan como una unidad. Contiene un archivo `go.mod` que especifica las dependencias del proyecto. Para inicializar un módulo Go:

1. **Inicializar el Módulo**:
   ```bash
   go mod init e42-go-challenge
   ```

   Esto crea un archivo `go.mod` en el directorio raíz del proyecto.

2. **Agregar Dependencias**:
   A medida que se escriben importaciones de paquetes externos en el código, Go las añadirá automáticamente al archivo `go.mod`.

### 3. Definir el Modelo de Asteroide

El modelo de asteroide se define en `internal/asteroid/model.go`:

```go
package asteroid

type Asteroid struct {
    ID               string  `json:"id"`
    Name             string  `json:"name"`
    Diameter         float64 `json:"diameter"`
    DiscoveryDate    string  `json:"discovery_date"`
    ObservationCount int     `json:"observation_count"`
    DistanceFromEarth float64 `json:"distance_from_earth"`
}
```

#### Explicación de Campos:

- **ID**: Identificador único del asteroide.
- **Name**: Nombre del asteroide.
- **Diameter**: Diámetro del asteroide.
- **DiscoveryDate**: Fecha en que se descubrió el asteroide.
- **ObservationCount**: Número de observaciones del asteroide.
- **DistanceFromEarth**: Distancia del asteroide a la Tierra.

#### 4. Aprendizajes:

  ```go
   import "go.mongodb.org/mongo-driver/bson/primitive"
   ```
   Esta librería ofrece tipos y funciones necesarios para trabajar con BSON (Binary JSON) en MongoDB. En particular, el tipo `primitive.ObjectID` se utiliza para representar los identificadores únicos (ObjectId) que MongoDB asigna a los documentos. Este tipo permite gestionar estos identificadores de manera adecuada dentro de la aplicación.

2. **¿Qué es esto: `json:"id" bson:"_id,omitempty"`?**
   Esta etiqueta se utiliza para especificar los nombres de los campos cuando el struct se serializa a JSON o BSON. Aquí está el desglose:
   - `json:"id"`: Indica que el campo `ID` debe ser serializado/deserializado con el nombre `id` en JSON.
   - `bson:"_id,omitempty"`: Indica que el campo `ID` debe ser serializado/deserializado con el nombre `_id` en BSON (formato utilizado por MongoDB). El atributo `omitempty` significa que el campo se omitirá en la serialización si está vacío o es su valor cero.

</details>

<details>
<summary><strong>3. Base de Datos</strong></summary>

- MongoDB
- Documentación Relevante

</details>

<details>
<summary><strong>4. Autenticación y Autorización</strong></summary>

- Protege tu API

</details>

<details>
<summary><strong>5. Testing</strong></summary>

- Pruebas Unitarias y de Integración

</details>

<details>
<summary><strong>6. Paginación y Filtros</strong></summary>

- Mejora de Usabilidad y Rendimiento

</details>

<details>
<summary><strong>7. Monitoring y Logging</strong></summary>

- Herramientas para Seguimiento en Tiempo Real

</details>

<details>
<summary><strong>8. Caching</strong></summary>

- Mejora de Rendimiento mediante Caché

</details>

<details>
<summary><strong>9. Frontend</strong></summary>

- Interfaz Web para Interactuar con la API

</details>

<details>
<summary><strong>10. Deployment y CI/CD</strong></summary>

- Configuración del Pipeline de CI/CD

</details>

<details>
<summary><strong>11. Escalabilidad</strong></summary>

- Preparación para Manejar Mayor Carga

</details>

<details>
<summary><strong>12. Mejoras en la Base de Datos</strong></summary>

- Optimización de la Base de Datos

</details>

<details>
<summary><strong>13. API Rate Limiting</strong></summary>

- Implementación de Limitación de Tasa para Proteger la API

</details>
