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

- Estructura de Archivos
  - **cmd/main.go**: Punto de entrada de la aplicación.
  - **internal/asteroid/**: Código relacionado con los asteroides.
    - **model.go**: Estructura (modelo) de los asteroides.
    - **repository.go**: Gestión del almacenamiento y recuperación de datos.
    - **handler.go**: Manejo de solicitudes HTTP relacionadas con los asteroides.
  - **internal/auth/**: Código relacionado con la autenticación y autorización.
  - **internal/router/**: Configuración de las rutas de la API.
- **Dockerfile**: Definición de la imagen Docker.
- **go.mod**: Archivo de configuración del módulo Go.

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
