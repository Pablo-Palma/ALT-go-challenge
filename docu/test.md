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

