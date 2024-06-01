# Documentación de Postman

## ¿Qué es Postman?

Postman es una herramienta poderosa para desarrollar, probar y documentar APIs. Permite a los desarrolladores enviar solicitudes HTTP a un servidor y ver las respuestas, lo que facilita la depuración y el desarrollo de APIs.

## Utilidades de Postman

1. **Desarrollo**: Facilita el desarrollo de APIs permitiendo probar endpoints rápidamente.
2. **Testing**: Permite escribir y ejecutar pruebas automatizadas para asegurar que la API funcione correctamente.
3. **Documentación**: Genera documentación interactiva que puede ser compartida con otros desarrolladores.
4. **Colaboración**: Facilita la colaboración entre equipos mediante la organización y compartición de colecciones de solicitudes.

## Beneficios de Usar Postman

- **Interfaz Intuitiva**: Fácil de usar, con una interfaz gráfica intuitiva.
- **Pruebas Automatizadas**: Permite crear pruebas automatizadas para validar el comportamiento de la API.
- **Documentación Automática**: Genera documentación de la API automáticamente.
- **Colaboración en Equipo**: Facilita la colaboración entre miembros del equipo al compartir colecciones de solicitudes.

## Instalación y Configuración de Postman

### 1. Descargar Postman

1. Visita la página de descarga de Postman:
   - [Postman Downloads](https://www.postman.com/downloads/)

2. Selecciona la versión para macOS.

### 2. Instalar Postman

1. Descarga el archivo `.zip`.
2. Extrae el archivo descargado.
3. Mueve la aplicación Postman a la carpeta `Aplicaciones`.

### 3. Iniciar Postman

- Abre Postman desde la carpeta `Aplicaciones`.

### 4. Crear una Cuenta o Iniciar Sesión

- Crea una cuenta gratuita o inicia sesión con una cuenta existente.

## Crear y Configurar una Colección en Postman

### Crear una Colección

1. Haz clic en "New" en la esquina superior izquierda.
2. Selecciona "Collection".
3. Nombra la colección, por ejemplo, "Asteroids API Tests".
4. Haz clic en "Create".

### Añadir Solicitudes a la Colección

#### Solicitud 1: Crear un Nuevo Asteroide

1. Añadir una nueva solicitud a la colección.
2. Nombra la solicitud, por ejemplo, "Create New Asteroid".
3. Configura la solicitud:
   - Método: `POST`
   - URL: `http://localhost:8080/api/v1/asteroides`
4. Añadir Encabezados:
   - Ve a la pestaña "Headers".
   - Añade un nuevo encabezado:
     - **Key**: `Content-Type`
     - **Value**: `application/json`
   - Añade otro encabezado para la autenticación:
     - **Key**: `Authorization`
     - **Value**: `Bearer <tu_token>`
5. Configurar el Cuerpo de la Solicitud:
   - Ve a la pestaña "Body".
   - Selecciona "raw".
   - Selecciona "JSON" en el menú desplegable al lado de "raw".
   - Introduce el cuerpo JSON de la solicitud:
     ```json
     {
       "name": "New Asteroid",
       "diameter": 150.5,
       "discovery_date": "01-01-2022",
       "observations": "Ninguna",
       "distances": []
     }
     ```
6. Haz clic en "Send" para enviar la solicitud y crear un nuevo asteroide.

#### Solicitud 2: Obtener Asteroides con Paginación

1. Añadir una nueva solicitud a la colección.
2. Nombra la solicitud, por ejemplo, "Get Asteroids with Pagination".
3. Configura la solicitud:
   - Método: `GET`
   - URL: `http://localhost:8080/api/v1/asteroides?page=1&limit=2`
4. Añadir Autorización:
   - Ve a la pestaña "Authorization".
   - Selecciona el tipo "Bearer Token".
   - Introduce tu token en el campo "Token".
5. Haz clic en "Send" para enviar la solicitud y ver la respuesta.

#### Solicitud 3: Obtener Asteroides con Paginación y Filtro por Nombre

1. Añadir una nueva solicitud a la colección.
2. Nombra la solicitud, por ejemplo, "Get Asteroids with Pagination and Name Filter".
3. Configura la solicitud:
   - Método: `GET`
   - URL: `http://localhost:8080/api/v1/asteroides?page=1&limit=2&name=Test`
4. Añadir Autorización como en la solicitud anterior.
5. Haz clic en "Send" para enviar la solicitud y ver la respuesta.

#### Solicitud 4: Obtener el Segundo Lote de Asteroides

1. Añadir una nueva solicitud a la colección.
2. Nombra la solicitud, por ejemplo, "Get Second Page of Asteroids".
3. Configura la solicitud:
   - Método: `GET`
   - URL: `http://localhost:8080/api/v1/asteroides?page=2&limit=2`
4. Añadir Autorización como en las solicitudes anteriores.
5. Haz clic en "Send" para enviar la solicitud y ver la respuesta.

## Respuestas Esperadas

### Solicitud para Crear un Nuevo Asteroide
Debe devolver un JSON con la información del asteroide creado, incluyendo su ID.

### Solicitud de Paginación Simple
Debe devolver un array JSON con los asteroides correspondientes a la primera página.

### Solicitud de Paginación con Filtro
Debe devolver un array JSON con los asteroides que coincidan con el nombre filtrado.

### Solicitud del Segundo Lote
Debe devolver un array JSON con los asteroides correspondientes a la segunda página.
