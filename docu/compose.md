### `compose.md`

## Docker Compose

### ¿Qué es Docker Compose?
Docker Compose es una herramienta que permite definir y ejecutar aplicaciones con múltiples contenedores Docker. Con Compose, puedes usar un archivo YAML para configurar los servicios de tu aplicación. Luego, con un solo comando, puedes crear e iniciar todos los servicios definidos en tu configuración.

### Beneficios de Docker Compose
- **Simplicidad**: Facilita la configuración y manejo de múltiples contenedores.
- **Consistencia**: Asegura que todos los desarrolladores y entornos usen la misma configuración.
- **Aislamiento**: Permite tener un entorno de desarrollo aislado con todas las dependencias necesarias.
- **Escalabilidad**: Facilita la gestión de aplicaciones que requieren múltiples servicios, como bases de datos, cachés y aplicaciones web.

### Por qué usar Docker Compose en nuestro proyecto
En nuestro proyecto, Docker Compose nos permite definir y gestionar fácilmente los contenedores de MongoDB y nuestra aplicación Go. Al usar Docker Compose, podemos asegurar que todos los desarrolladores tengan un entorno de desarrollo consistente y aislado, simplificando la configuración y despliegue de nuestra aplicación.

### Guía para levantar la API

#### 1. Construir los contenedores
```bash
docker-compose build
```

#### 2. Levantar los contenedores
```bash
docker-compose up
```

#### 3. Detener y eliminar los contenedores
```bash
docker-compose down
```

### docker-compose.yaml explanations and conventions

#### Version
```yaml
version: '3.8'
```
La versión especifica la versión del formato de Docker Compose que se está utilizando. En este caso, '3.8' se refiere a una versión específica del archivo de configuración de Docker Compose. Las versiones definen qué características de Compose están disponibles y cómo se deben interpretar las configuraciones.

#### Services
```yaml
services:
  mongodb:
    ...
  app:
    ...
```
`services` es la sección donde se definen los contenedores que componen la aplicación. Cada servicio se ejecuta en un contenedor separado y puede comunicarse con otros servicios según lo definido.

##### MongoDB Service
```yaml
mongodb:
  image: mongo:latest
  container_name: mongodb
  ports:
    - "27017:27017"
  volumes:
    - mongo-data:/data/db
```
- `image`: Define la imagen de Docker que se usará para el contenedor. En este caso, se usa la última versión de MongoDB.
- `container_name`: Asigna un nombre específico al contenedor para facilitar su identificación.
- `ports`: Expone el puerto 27017 del contenedor para que sea accesible desde el host.
- `volumes`: Monta un volumen para persistir los datos de la base de datos.

##### App Service
```yaml
app:
  build:
    context: .
  container_name: go-app
  ports:
    - "8080:8080"
  environment:
    - MONGO_URI=mongodb://mongodb:27017/asteroidsdb
  depends_on:
    - mongodb
```
- `build`: Define que el contenedor debe construirse desde el contexto actual (directorio actual).
- `container_name`: Asigna un nombre específico al contenedor.
- `ports`: Expone el puerto 8080 del contenedor para que sea accesible desde el host.
- `environment`: Establece variables de entorno necesarias para la configuración de la aplicación.
- `depends_on`: Define que este servicio depende de que el servicio `mongodb` esté activo y funcionando.

#### Volumes
```yaml
volumes:
  mongo-data:
```
Los volúmenes se utilizan para persistir datos fuera del ciclo de vida de los contenedores, asegurando que los datos no se pierdan cuando se detienen o eliminan los contenedores. En este caso, `mongo-data` se utiliza para almacenar los datos de MongoDB.

### Conventions
- **Indentation**: La indentación en YAML es importante y se hace con dos espacios para cada nivel de profundidad.
- **Hierarchy**: Cada nivel de indentación representa una jerarquía. Por ejemplo, `services` es el nivel principal, `mongodb` y `app` son subniveles, y sus configuraciones específicas son aún más profundas.
- **Key-Value Pairs**: La sintaxis se basa en pares clave-valor, con `:` separando las claves de sus valores.
