## Inicialización de la API con Docker

### 1. Construir la Imagen Docker

Construye la imagen Docker de tu aplicación:

```sh
docker build -t alt-go-challenge .
```

- **`docker build`**: Comando para construir una imagen Docker.
- **`-t alt-go-challenge`**: Asigna una etiqueta (nombre) a la imagen que estás construyendo.
- **`.`**: Indica que Dockerfile y el contexto de construcción están en el directorio actual.

### 2. Iniciar el Contenedor de MongoDB

Inicia el contenedor de MongoDB:

```sh
docker run -d --name mongodb -p 27017:27017 mongo
```

- **`docker run`**: Comando para crear y ejecutar un contenedor.
- **`-d`**: Ejecuta el contenedor en modo desacoplado (detached), es decir, en segundo plano.
- **`--name mongodb`**: Asigna un nombre al contenedor (`mongodb`).
- **`-p 27017:27017`**: Mapea el puerto 27017 del contenedor al puerto 27017 del host, permitiendo el acceso a MongoDB.
- **`mongo`**: Especifica la imagen de Docker a utilizar (en este caso, la imagen oficial de MongoDB).

### 3. Iniciar el Contenedor de la API

Inicia el contenedor de tu aplicación API:

```sh
docker run -d --name alt-go-challenge -p 8080:8080 --link mongodb:mongo alt-go-challenge
```

- **`docker run`**: Comando para crear y ejecutar un contenedor.
- **`-d`**: Ejecuta el contenedor en modo desacoplado (detached), es decir, en segundo plano.
- **`--name alt-go-challenge`**: Asigna un nombre al contenedor (`alt-go-challenge`).
- **`-p 8080:8080`**: Mapea el puerto 8080 del contenedor al puerto 8080 del host, permitiendo el acceso a la API.
- **`--link mongodb:mongo`**: Conecta el contenedor `alt-go-challenge` al contenedor `mongodb`, utilizando el alias `mongo` para resolver el nombre del contenedor MongoDB.
- **`alt-go-challenge`**: Especifica la imagen de Docker a utilizar (la imagen construida en el paso 1).

### Resumen de Comandos

```sh
# Construir la imagen Docker de la aplicación
docker build -t alt-go-challenge .

# Iniciar el contenedor de MongoDB
docker run -d --name mongodb -p 27017:27017 mongo

# Iniciar el contenedor de la API
docker run -d --name alt-go-challenge -p 8080:8080 --link mongodb:mongo alt-go-challenge
```

