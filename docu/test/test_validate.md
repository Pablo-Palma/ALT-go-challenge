### Testing Validation 

1. **Intentar crear un asteroide con nombre vacío**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid name`.

2. **Intentar crear un asteroide con nombre lleno**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `201 Created` con la información del asteroide creado.

3. **Intentar crear un asteroide con diámetro inválido (negativo)**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": -123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid diameter`.

4. **Intentar crear un asteroide con diámetro válido**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `201 Created` con la información del asteroide creado.

5. **Intentar crear un asteroide con formato de fecha inválido**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "2024-01-01", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid date format`.

6. **Intentar crear un asteroide con formato de fecha válido**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": []}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `201 Created` con la información del asteroide creado.

7. **Intentar crear un asteroide con distancia inválida (negativa)**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": [{"date": "01-01-2024", "distance": -1}]}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid distance`.

8. **Intentar crear un asteroide con distancia válida**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/asteroides -H "Content-Type: application/json" -d '{"name": "Test Asteroid", "diameter": 123.4, "discovery_date": "01-01-2024", "observations": "Ninguna", "distances": [{"date": "01-01-2024", "distance": 1.2}]}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `201 Created` con la información del asteroide creado.

### Pruebas de Validación de Actualización

1. **Intentar actualizar un asteroide con nombre vacío**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"name": ""}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid name`.

2. **Intentar actualizar un asteroide con nombre lleno**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"name": "Updated Test Asteroid"}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `200 OK` con la información del asteroide actualizado.

3. **Intentar actualizar un asteroide con diámetro inválido (negativo)**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"diameter": -123.4}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid diameter`.

4. **Intentar actualizar un asteroide con diámetro válido**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"diameter": 123.4}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `200 OK` con la información del asteroide actualizado.

5. **Intentar actualizar un asteroide con formato de fecha inválido**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"discovery_date": "2024-01-01"}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid date format`.

6. **Intentar actualizar un asteroide con formato de fecha válido**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"discovery_date": "01-01-2024"}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `200 OK` con la información del asteroide actualizado.

7. **Intentar actualizar un asteroide con distancia inválida (negativa)**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"distances": [{"date": "01-01-2024", "distance": -1}]}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `400 Bad Request` con mensaje `invalid distance`.

8. **Intentar actualizar un asteroide con distancia válida**:
   ```bash
   curl -X PATCH http://localhost:8080/api/v1/asteroides/{id} -H "Content-Type: application/json" -d '{"distances": [{"date": "01-01-2024", "distance": 1.2}]}' -b cookies.txt
   ```
   **Esperado**: Deberías recibir `200 OK` con la información del asteroide actualizado.

### Verificación de Resultados

#### Creación

- **Nombre vacío**: Deberías recibir `400 Bad Request` con mensaje `invalid name`.
- **Nombre lleno**: Deberías recibir `201 Created` con la información del asteroide creado.
- **Diámetro inválido**: Deberías recibir `400 Bad Request` con mensaje `invalid diameter`.
- **Diámetro válido**: Deberías recibir `201 Created` con la información del asteroide creado.
- **Fecha inválida**: Deberías recibir `400 Bad Request` con mensaje `invalid date format`.
- **Fecha válida**: Deberías recibir `201 Created` con la información del asteroide creado.
- **Distancia inválida**: Deberías recibir `400 Bad Request` con mensaje `invalid distance`.
- **Distancia válida**: Deberías recibir `201 Created` con la información del asteroide creado.

#### Actualización

- **Nombre vacío**: Deberías recibir `400 Bad Request` con mensaje `invalid name`.
- **Nombre lleno**: Deberías recibir `

200 OK` con la información del asteroide actualizado.
- **Diámetro inválido**: Deberías recibir `400 Bad Request` con mensaje `invalid diameter`.
- **Diámetro válido**: Deberías recibir `200 OK` con la información del asteroide actualizado.
- **Fecha inválida**: Deberías recibir `400 Bad Request` con mensaje `invalid date format`.
- **Fecha válida**: Deberías recibir `200 OK` con la información del asteroide actualizado.
- **Distancia inválida**: Deberías recibir `400 Bad Request` con mensaje `invalid distance`.
- **Distancia válida**: Deberías recibir `200 OK` con la información del asteroide actualizado.
