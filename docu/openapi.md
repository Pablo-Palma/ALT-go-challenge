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
