## Documentación Auth


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

