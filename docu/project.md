# API RESTful in Go Language

Este proyecto creamos una **API RESTful**, un tipo de API que sigue los principios del estilo de arquitectura REST (Representational Sta    te Transfer). REST es una forma de diseñar APIs que utiliza HTTP para realizar operaciones CRUD (Crear, Leer    , Actualizar, Eliminar) sobre los recursos representados en formato JSON o XML.

Elegimos GO como  lenguaje de programación  para desplegar una aplicación contenerizada con **docker**, una aplicaición de código abierto que nos permite automatizar el despligue de aplicaciones dentro de contenedores de software.

Nuestra API almacena los datos tanto de los usuarios como de los objetos en una base de datos NoSQL, en este caso hemos elegido **mongodb**.

### Metodos CRUD(Create, Read, Update, Delete)

* **Almacenar información** sobre asteroides: incluyendo nombre, tamaño, composición, etc.
* **Registrar la distancia** del asteroide a la Tierra en diferentes fechas.
* Permitir la **consulta de información** sobre asteroides específicos.
* Permitir la **actualización de la información** de un asteroide.
* **Eliminar información** de asteroides que ya no son una amenaza.

### Funciones adicionales:

* **Paginación y filtros**.
* **Sistema de autenticación y autorización** para controlar el acceso a la API.
* **Validación de entrada de datos** para evitar errores y datos incorrectos.
* **Monitoring:** Uso de herramientas de trazas, logging y Monitorizacion de codigos de estado.
* **Documentación:** Hemos empleado OpenApi para describir, producir, consumir y visualizar servicios web RESTful.
* **GitHub** como repositorio de almacenamiento y control de versiones.
* **Testing:** Hemos incluido test unitarios, tanto en go como postman.
* **Frontend.** Usamos Astro para renderizar la aplicación en el navegador.

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

# Documentación

<details>
<summary><strong>1. OpenApi</strong></summary>

# Documentación de OpenAPI

## ¿Qué es OpenAPI?

OpenAPI es una especificación estándar para definir APIs RESTful. Permite a los desarrolladores describir todos los aspectos de su API, incluyendo endpoints, métodos, parámetros, respuestas y más, en un formato legible tanto por humanos como por máquinas, generalmente utilizando YAML o JSON.

### ¿Qué es una API RESTful?

Una API RESTful es una interfaz de programación de aplicaciones (API) que sigue los principios de REST (Representational State Transfer). REST es un estilo de arquitectura para diseñar servicios web. Las APIs RESTful permiten la comunicación entre sistemas a través de HTTP usando métodos estándar como GET, POST, PUT, DELETE, entre otros.

### Características de una API RESTful

- **Stateless**: Cada solicitud del cliente al servidor debe contener toda la información necesaria para entender y procesar la solicitud. El servidor no guarda estado entre las solicitudes.
- **Cacheable**: Las respuestas deben definir si son cacheables o no para mejorar el rendimiento.
- **Uniform Interface**: Utiliza una interfaz uniforme para que las interacciones entre el cliente y el servidor sean simples y coherentes.
- **Client-Server**: La arquitectura está dividida en cliente y servidor, lo que permite una separación de preocupaciones.

### Utilidades de OpenAPI

1. **Documentación**: Proporciona documentación clara y precisa de la API.
2. **Generación de Código**: Permite generar código de cliente y servidor en varios lenguajes.
3. **Validación**: Facilita la validación de las especificaciones de la API.
4. **Interoperabilidad**: Facilita la integración y colaboración entre equipos y sistemas.
5. **Testing**: Mejora el proceso de pruebas automatizadas.

### Beneficios de OpenAPI

- **Consistencia**: Asegura que la API sea consistente y siga estándares definidos.
- **Productividad**: Aumenta la productividad al permitir la generación automática de documentación y código.
- **Mantenimiento**: Facilita el mantenimiento de la API, ya que cualquier cambio en la especificación se refleja automáticamente en la documentación y en el código generado.
- **Comunicación**: Mejora la comunicación entre los equipos de desarrollo, QA y operaciones, proporcionando una fuente de verdad única y comprensible.

## Comandos Útiles y Flags

### Ejecución de Swagger UI con Docker

El siguiente comando ejecuta Swagger UI utilizando Docker, permitiendo visualizar la documentación de tu API definida en OpenAPI.

```bash
docker run -p 8080:8080 -e SWAGGER_JSON=/foo/openapi.yaml -v $(pwd)/openapi.yaml:/foo/openapi.yaml swaggerapi/swagger-ui
```

#### Flags Explicados

- `-p 8080:8080`: Este flag mapea el puerto 8080 del contenedor al puerto 8080 de la máquina host. La `-p` viene de `--publish` y se usa para exponer puertos.
- `-e SWAGGER_JSON=/foo/openapi.yaml`: Este flag establece una variable de entorno (`-e` viene de `--env`) dentro del contenedor. `SWAGGER_JSON` es la variable que Swagger UI utiliza para saber la ubicación del archivo de especificación OpenAPI.
- `-v $(pwd)/openapi.yaml:/foo/openapi.yaml`: Este flag monta un volumen (`-v` viene de `--volume`) de la máquina host al contenedor, permitiendo que el contenedor acceda al archivo `openapi.yaml` en el host.

### Validación de la Especificación OpenAPI

El siguiente comando valida un archivo de especificación OpenAPI para asegurarse de que cumple con la especificación estándar.

```bash
swagger-cli validate openapi.yaml
```

#### Flags Explicados

- `validate`: Este subcomando de `swagger-cli` se usa para validar el archivo de especificación.
- `openapi.yaml`: Es el archivo de especificación que se desea validar. `swagger-cli` verificará que este archivo cumpla con la sintaxis y las reglas de la especificación OpenAPI.

## Uso de OpenAPI en tu Proyecto

1. **Crear la Especificación**: Define tu API en un archivo YAML o JSON siguiendo la especificación OpenAPI.
2. **Validar la Especificación**: Usa `swagger-cli validate openapi.yaml` para asegurarte de que tu especificación es válida.
3. **Generar Código**: Opcionalmente, usa herramientas como `oapi-codegen` para generar el código de servidor y cliente a partir de tu especificación.
4. **Documentar la API**: Ejecuta Swagger UI usando Docker para visualizar la documentación de tu API y facilitar su uso por parte de otros desarrolladores.

</details>

## Implementación de la Aplicación en Go

<details>
<summary><strong>2. Asteroids </strong></summary>

### 1. Estructura:

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

### 4. Aprendizajes:

**Serialización:** Convertir una estructura de datos a un formato que puede ser almacenado o transmitido.

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
<summary><strong>3. Autenticación y Autorización </strong></summary>

## Documentación Auth

- **Registro**: Genera un hash de la contraseña y almacena el usuario.
- **Inicio de Sesión**: Verifica las credenciales, genera un JWT y lo almacena en una cookie con duración de 24 horas.
- **Middleware**: Verifica el token JWT en la cookie o en el encabezado de autorización para proteger los endpoints.
- **Solicitud con `curl`**: El comando `curl` que proporcionaste utiliza cookies para enviar el token JWT almacenado en `cookies.txt`.

El comando `curl` en tu ejemplo **usa cookies** para la autenticación, ya que incluye la opción `-b cookies.txt` para enviar la cookie con el token JWT almacenado en dicho archivo.

### Uso de Tokens y Cookies:
Tu implementación permite el uso de tokens tanto en cookies como en el encabezado de autorización. Aquí está el flujo detallado de cómo funciona:

1. **Login y Set-Cookie**:
    - Después de un inicio de sesión exitoso, el token JWT se envía al cliente en una cookie HTTP.
    - Ejemplo:
      ```go
      http.SetCookie(w, &http.Cookie{
          Name:    "token",
          Value:   token,
          Expires: time.Now().Add(24 * time.Hour),
      })
      ```

2. **Auth Middleware**:
    - Este middleware intercepta las solicitudes a los endpoints protegidos.
    - Busca el token JWT en las cookies o en el encabezado de autorización.
    - Si el token es válido, extrae el nombre de usuario de las reclamaciones (`claims`) y lo almacena en el contexto de la solicitud.


### Resumen:
- **Autenticación**: Se realiza mediante nombres de usuario y contraseñas, verificadas con `bcrypt`. Si es correcta, se genera un token JWT.
- **Autorización**: Se realiza mediante el middleware que verifica el token JWT en cada solicitud protegida.
- **Uso de Tokens**: Los tokens JWT se pueden pasar en cookies o en el encabezado de autorización.

Este enfoque proporciona flexibilidad en la forma en que los clientes pueden enviar tokens, y garantiza que solo los usuarios autenticados puedan acceder a los recursos protegidos.

## Headers HTTP

### ¿Qué es un Encabezado HTTP?
Un encabezado HTTP es una parte de la solicitud o respuesta HTTP que proporciona metadatos adicionales sobre la transacción HTTP. Los encabezados son pares clave-valor enviados en el formato:
```
Clave: Valor
```

### Tipos Comunes de Encabezados HTTP

#### Encabezados de Solicitud (Request Headers)
- **Authorization**: Incluye credenciales para autenticar al cliente con el servidor.
- **Content-Type**: Indica el tipo de contenido que se está enviando al servidor (e.g., `application/json`).
- **Accept**: Indica qué tipo de respuesta espera el cliente (e.g., `application/json`).

#### Encabezados de Respuesta (Response Headers)
- **Content-Type**: Indica el tipo de contenido de la respuesta del servidor.
- **Set-Cookie**: Establece cookies que el navegador almacenará y enviará en futuras solicitudes.

### Encabezado de Autorización

#### ¿Qué es un Token en el Encabezado de Autorización?
Un token es una cadena de caracteres utilizada para identificar a un usuario autenticado. Los tokens se pueden enviar en el encabezado `Authorization` para autenticar solicitudes HTTP. Un esquema común es el Bearer, utilizado para tokens JWT (JSON Web Tokens).

#### Formato del Encabezado de Autorización
El token se envía en el encabezado `Authorization` de la solicitud HTTP utilizando el esquema Bearer:
```
Authorization: Bearer <token_jwt>
```

### Uso de Tokens en el Desarrollo Web
- **Autenticación de API**: Los tokens se utilizan para autenticar solicitudes a APIs RESTful. Al incluir el token en el encabezado `Authorization`, el servidor puede verificar la identidad del cliente y autorizar la solicitud.
- **Seguridad**: Los tokens son más seguros que las cookies porque no se almacenan en el navegador. Además, los tokens pueden tener un tiempo de expiración, limitando su validez temporal.
- **Estándar en el Desarrollo Web**: Usar tokens en el encabezado `Authorization` es una práctica estándar en el desarrollo web moderno, especialmente para APIs.

## JWT

**JWT:** (JSON Web Tokens)
**Claims:** Declaraciones sobre una entidad, (normalmente, el usuario) y datos adicionales.
En el contexto JWT, son parte del token y se utilizan para transmitir información entre el servidor y el cliente de forma segura.
Las claims puede ser estándar (como `iss` para el emisor, `exp` para la expiración) o personalizadas (como `username` en nuestro caso).

### Autenticación y Autorización en la API de Registro de Asteroides

La API de Registro de Asteroides implementa autenticación y autorización mediante JWT (JSON Web Tokens). A continuación se explica cómo funciona y cómo puedes interactuar con ella.

### Cómo Funciona

1. **Registro de Usuario**: Los usuarios se registran proporcionando un nombre de usuario y una contraseña. La contraseña se almacena de manera segura mediante un hash.
2. **Inicio de Sesión**: Los usuarios inician sesión proporcionando sus credenciales. Si las credenciales son válidas, se genera un token JWT y se envía al cliente.
3. **Acceso a Endpoints Protegidos**: Para acceder a ciertos endpoints, el cliente debe enviar el token JWT como parte de la solicitud. El servidor verifica el token para autorizar el acceso.

### Comandos para Interactuar con la API

#### 1. Registrar un Nuevo Usuario

Registra un nuevo usuario enviando una solicitud `POST` al endpoint `/register` con un cuerpo JSON que contiene el `username` y `password`.

```bash
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}'
```

#### 2. Iniciar Sesión con el Usuario Registrado

Inicia sesión enviando una solicitud `POST` al endpoint `/login` con un cuerpo JSON que contiene el `username` y `password`. La opción `-c cookies.txt` guarda las cookies recibidas (incluyendo el token JWT) en un archivo llamado `cookies.txt`.

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}' -c cookies.txt
```

#### 3. Acceder a un Endpoint Protegido

Accede a un endpoint protegido enviando una solicitud `GET` al endpoint `/protected` usando las cookies guardadas en `cookies.txt` (que contienen el token JWT).

```bash
curl -X GET http://localhost:8080/protected -b cookies.txt
```

### Explicación de los Comandos

- **Registro**: El primer comando registra un nuevo usuario enviando una solicitud `POST` al endpoint `/register` con un cuerpo JSON que contiene el `username` y `password`.
- **Inicio de Sesión**: El segundo comando inicia sesión con el usuario registrado enviando una solicitud `POST` al endpoint `/login` con un cuerpo JSON que contiene el `username` y `password`. La opción `-c cookies.txt` guarda las cookies recibidas (incluyendo el token JWT) en un archivo llamado `cookies.txt`.
- **Acceso al Endpoint Protegido**: El tercer comando accede al endpoint protegido enviando una solicitud `GET` al endpoint `/protected` usando las cookies guardadas en `cookies.txt` (que contienen el token JWT).

### Verificación y Resolución de Problemas

- **Error 401 Unauthorized**: Asegúrate de que el usuario está registrado correctamente y que el token JWT está siendo enviado correctamente en la cookie.
- **Error 404 Not Found**: Asegúrate de que los endpoints `/register`, `/login`, y `/protected` están definidos correctamente en tu código.
- **Errores en los Logs**: Verifica los logs de Docker para obtener más información sobre cualquier error que ocurra en tu aplicación.

</details>

## Base de datos

<details>
<summary><strong>4. Mongodb </strong></summary>

  ### ¿Qué es MongoDB?

MongoDB es una base de datos NoSQL de código abierto, orientada a documentos, que almacena datos en formato BSON (Binary JSON). A diferencia de las bases de datos relacionales que usan tablas y filas, MongoDB usa colecciones y documentos, permitiendo una estructura de datos más flexible y escalable.

### Beneficios de Usar MongoDB

1. **Flexibilidad**: Permite almacenar datos en un formato JSON-like, lo que facilita la manipulación y consulta de datos complejos.
2. **Escalabilidad**: Soporta escalado horizontal a través de sharding, lo que permite manejar grandes volúmenes de datos y tráfico.
3. **Alta Disponibilidad**: Ofrece replicación automática y recuperación ante fallos con configuraciones de replica sets.
4. **Rendimiento**: Optimizado para manejar grandes cantidades de datos y operaciones de lectura/escritura de forma eficiente.
5. **Desarrollo Ágil**: La flexibilidad de su esquema permite cambios rápidos en la estructura de los datos sin necesidad de grandes migraciones.

### Comparativa: SQL vs NoSQL

| Característica           | SQL (Relacional)                      | NoSQL (MongoDB)                            |
|--------------------------|---------------------------------------|--------------------------------------------|
| **Modelo de Datos**      | Tablas y filas                        | Documentos, clave-valor, columnas o grafos |
| **Esquema**              | Estructura fija (esquema rígido)      | Estructura flexible (esquema dinámico)     |
| **Escalabilidad**        | Vertical (más potente el hardware)    | Horizontal (más servidores)                |
| **Transacciones**        | Soporte completo ACID                 | Soporte parcial ACID                       |
| **Consultas**            | Lenguaje SQL (consultas complejas)    | Consultas flexibles con JSON               |
| **Casos de Uso**         | Sistemas transaccionales, ERP, CRM    | Big Data, análisis en tiempo

 real, IoT     |
| **Ejemplos**             | MySQL, PostgreSQL, Oracle             | MongoDB, Cassandra, Couchbase              |

### Resumen de Comandos

1. **Instalar MongoDB y mongosh**:

   ```sh
   brew tap mongodb/brew
   brew install mongodb-community
   brew install mongosh
   ```

2. **Iniciar MongoDB**:

   ```sh
   brew services start mongodb/brew/mongodb-community
   ```

3. **Acceder a MongoDB Shell y Configurar la Base de Datos**:

   ```sh
   mongosh
   use asteroidsdb
   db.createCollection("asteroids")
   ```

4. **Ejecución de MongoDB en Docker**:

   ```sh
   docker run -d --name mongodb -p 27017:27017 mongo
   mongosh "mongodb://localhost:27017"
   use asteroidsdb
   ```

### Comandos para Verificar la Base de Datos y Colecciones

1. **Acceder a MongoDB en el contenedor Docker**:

   ```sh
   docker exec -it mongodb mongosh
   ```

2. **Mostrar todas las bases de datos**:

   ```sh
   show dbs
   ```

3. **Seleccionar la base de datos `asteroidsdb`**:

   ```sh
   use asteroidsdb
   ```

4. **Mostrar todas las colecciones en `asteroidsdb`**:

   ```sh
   show collections
   ```

5. **Listar documentos en la colección `asteroids`**:

   ```sh
   db.asteroids.find().pretty()
   ```

6. **Listar documentos en la colección `users`**:

   ```sh
   db.users.find().pretty()
   ```

</details>

## Logging and monitoring

<details>
<summary><strong>5. Prometheus </strong></summary>

# Configuración de Prometheus para la API de Registro de Asteroides

## Introducción a Prometheus

**Prometheus** es una herramienta de monitoreo y alerta de código abierto que recolecta y almacena métricas como datos de series temporales. Es ampliamente utilizada para monitorear aplicaciones y sistemas.

### Beneficios de Prometheus

- **Recolecta métricas de series temporales.**
- **Permite consultas avanzadas con PromQL.**
- **Soporta alertas basadas en métricas.**
- **Funciona de manera independiente y se integra con múltiples sistemas.**

## Configuración de Prometheus

### Acceder a Prometheus

Abre Prometheus en tu navegador web en `http://localhost:9090`.

### Configuración de Prometheus (`prometheus.yml`)

El archivo de configuración de Prometheus define los targets que debe monitorear. Aquí está el ejemplo actual:

```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
    - targets: ['localhost:9090']

  - job_name: 'go-app'
    static_configs:
    - targets: ['go-app:8080']
```

### Consultas Básicas en Prometheus

#### Total de Solicitudes HTTP por Endpoint

```promql
sum by (endpoint) (http_requests_total)
```

#### Tasa de Solicitudes HTTP por Segundo para un Endpoint Específico

```promql
rate(http_requests_total{endpoint="/api/v1/asteroides"}[1m])
```

#### Uso de Memoria

```promql
go_memstats_alloc_bytes
```

#### Número de Goroutines

```promql
go_goroutines
```

## Integración en el Proyecto

La integración de Prometheus en el proyecto se refleja en varios archivos:

1. **`docker-compose.yml`**: Define los servicios de Docker, incluyendo Prometheus, Grafana y la aplicación Go.
2. **`prometheus.yml`**: Configura los targets que Prometheus debe monitorear.
3. **`main.go`**: Implementa las métricas de Prometheus en la aplicación Go.

### Configuración del Servicio en `docker-compose.yml`

```yaml
prometheus:
  image: prom/prometheus
  container_name: prometheus
  ports:
    - "9090:9090"
  volumes:
    - ./prometheus:/etc/prometheus
  command:
    - '--config.file=/etc/prometheus/prometheus.yml'
```

### Implementación de Métricas en `main.go`

Ejemplo de cómo se exponen las métricas en la aplicación Go:

```go
import (
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "net/http"
)

var httpRequestsTotal = prometheus.NewCounterVec(
  prometheus.CounterOpts{
    Name: "http_requests_total",
    Help: "Total number of HTTP requests",
  },
  []string{"method", "endpoint"},
)

func init() {
  prometheus.MustRegister(httpRequestsTotal)
}

func prometheusMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
    next.ServeHTTP(w, r)
  })
}

func main() {
  http.Handle("/metrics", promhttp.Handler())
  http.ListenAndServe(":8080", nil)
}
```


</details>

<details>
<summary><strong>6. Grafana</strong></summary>

# Dashboard de Métricas de la API de Registro de Asteroides

## Introducción a Grafana

**Grafana** es una plataforma de código abierto para la visualización y monitoreo de métricas, que permite a los usuarios consultar, visualizar y alertar sobre datos de series temporales. Es ampliamente utilizada para observar la infraestructura y aplicaciones de TI. Grafana soporta múltiples fuentes de datos, incluyendo Prometheus, InfluxDB, Graphite, Elasticsearch, MySQL, y muchos más.

### Beneficios de Grafana

- **Visualización**: Proporciona una interfaz gráfica rica para crear dashboards personalizados que muestran métricas y datos en tiempo real.
- **Alertas**: Permite configurar alertas basadas en las métricas monitoreadas para notificar sobre eventos críticos.
- **Integración**: Se integra con múltiples fuentes de datos, lo que facilita la recopilación y visualización de métricas de diferentes sistemas.
- **Facilidad de uso**: Ofrece una interfaz intuitiva que facilita la creación y gestión de dashboards sin necesidad de conocimientos avanzados en programación.

### Integración en el Proyecto

La integración de Grafana en el proyecto se refleja en varios archivos de configuración y documentación:

1. **Archivo `docker-compose.yml`**: Define los servicios de Docker, incluyendo Grafana, Prometheus y la aplicación Go.
2. **Archivo de Configuración de Prometheus (`prometheus.yml`)**: Configura los targets que Prometheus debe monitorear.
3. **Código de la Aplicación (`main.go`)**: Implementa las métricas de Prometheus en la aplicación Go.
4. **Documentación (`grafana.md`)**: Detalla la configuración y uso de dashboards en Grafana.

### Configuración de Grafana

#### Acceder a Grafana

Abre Grafana en tu navegador web en `http://localhost:3000` y accede con las credenciales configuradas (por defecto, admin/admin).

#### Configurar Prometheus como Fuente de Datos

1. Ve a "Configuration" > "Data Sources".
2. Añade una nueva fuente de datos y selecciona "Prometheus".
3. Configura la URL de Prometheus (por ejemplo, `http://prometheus:9090`) y guarda la configuración.

## Panel de Métricas HTTP

**Nombre del Panel:** Métricas HTTP

**Descripción del Panel:** Este panel muestra las métricas de solicitudes HTTP para la API de registro de asteroides, incluyendo el total de solicitudes por endpoint, la tasa de solicitudes por segundo y la distribución de solicitudes por método.

**Consultas y Tipos de Visualización:**

- **Total de Solicitudes HTTP por Endpoint:**
  - **Query:**
    ```promql
	sum by (endpoint) (rate(http_requests_total[1m]))
    ```
  - **Tipo de Visualización:** Bar chart

- **Tasa de Solicitudes HTTP por Segundo para un Endpoint Específico:**
  - **Query:**
    ```promql
    rate(http_requests_total{endpoint="/api/v1/asteroides"}[1m])
    ```
  - **Tipo de Visualización:** Line chart

- **Número Total de Solicitudes HTTP Agrupadas por Método:**
  - **Query:**
    ```promql
	sum by (method) (rate(http_requests_total[1m]))
    ```
  - **Tipo de Visualización:** Bar chart

## Panel de Métricas de Rendimiento

**Nombre del Panel:** Métricas de Rendimiento

**Descripción del Panel:** Este panel muestra las métricas de rendimiento para la API de registro de asteroides, incluyendo el uso de memoria, el número de goroutines en ejecución y los tiempos de respuesta HTTP.

**Consultas y Tipos de Visualización:**

- **Uso de Memoria:**
  - **Query:**
    ```promql
    go_memstats_alloc_bytes
    ```
  - **Tipo de Visualización:** Line chart

- **Número de Goroutines:**
  - **Query:**
    ```promql
    go_goroutines
    ```
  - **Tipo de Visualización:** Line chart

- **Tiempos de Respuesta HTTP (si tienes esta métrica implementada):**
  - **Query:**
    ```promql
    histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le))
    ```
  - **Tipo de Visualización:** Line chart



</details>

## Bonus

<details>
<summary><strong>7. Paginación y Filtros</strong></summary>

- Herramientas para Seguimiento en Tiempo Real

</details>

<details>
<summary><strong>8. Frontend</strong></summary>

- Interfaz Web para Interactuar con la API

</details>

<details>
<summary><strong>9. Testing</strong></summary>

- Pruebas Unitarias y de Integración

</details>
