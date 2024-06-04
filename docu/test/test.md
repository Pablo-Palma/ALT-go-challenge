### Testing AUTH in proccess

1. **Acceso no autorizado**:
   ```bash
   curl -X GET http://localhost:8080/protected
   ```

2. **Registrar usuario**:
   ```bash
   curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}'
   ```

3. **Iniciar sesión**:
   ```bash
   curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}' -c cookies.txt
   ```

4. **Eliminar sin autenticación**:
   ```bash
   curl -X DELETE "http://localhost:8080/deleteuser?username=testuser"
   ```

5. **Eliminar con autenticación**:
   ```bash
   curl -X DELETE "http://localhost:8080/deleteuser?username=testuser" -b cookies.txt
   ```

### Verificación de Resultados

- **Acceso no autorizado**: Deberías recibir `401 Unauthorized`.
- **Registro**: Deberías recibir una respuesta con la información del usuario registrado.
- **Inicio de sesión**: Deberías recibir un token JWT en las cookies.
- **Eliminar sin autenticación**: Deberías recibir `401 Unauthorized`.
- **Eliminar con autenticación**: Deberías recibir un mensaje indicando que el usuario fue eliminado exitosamente.


### Testing permission on Asteroid methods (CUD)

1. **Intentar crear un asteroide sin autenticación**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2020", "observations": "Ninguna", "distances": []}'
   ```

2. **Intentar actualizar el asteroide sin autenticación** (usa el ID que intentaste crear anteriormente):
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"name": "Updated Test Asteroid"}'
   ```

3. **Intentar eliminar el asteroide sin autenticación** (usa el mismo ID):
   ```bash
   curl -X DELETE http://localhost:8080/api/v1/asteroides/{id}
   ```

4. **Registrar usuario**:
   ```bash
   curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}'
   ```

5. **Iniciar sesión**:
   ```bash
   curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testuser", "password": "password"}' -c cookies.txt
   ```

6. **Crear un asteroide con autenticación**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2020", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```

7. **Actualizar el asteroide con autenticación** (usa el ID que recibiste en la respuesta de creación):
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"name": "Updated Test Asteroid"}' -b cookies.txt
   ```

8. **Eliminar el asteroide con autenticación** (usa el mismo ID):
   ```bash
   curl -X DELETE http://localhost:8080/api/v1/asteroides/{id} -b cookies.txt
   ```

### Verificación de Resultados

- **Intentar crear un asteroide sin autenticación**: Deberías recibir `401 Unauthorized`.
- **Intentar actualizar el asteroide sin autenticación**: Deberías recibir `401 Unauthorized`.
- **Intentar eliminar el asteroide sin autenticación**: Deberías recibir `401 Unauthorized`.
- **Registro**: Deberías recibir una respuesta con la información del usuario registrado.
- **Inicio de sesión**: Deberías recibir una cookie con el token JWT.
- **Crear un asteroide con autenticación**: Deberías recibir una respuesta con la información del asteroide creado, incluyendo su ID.
- **Actualizar el asteroide con autenticación**: Deberías recibir una respuesta con la información del asteroide actualizado.
- **Eliminar el asteroide con autenticación**: Deberías recibir una respuesta indicando que el asteroide fue eliminado exitosamente.
