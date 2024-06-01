### Paginación y Filtros

## ¿Qué es la Paginación y Filtros?

**Paginación** y **filtros** son técnicas utilizadas para mejorar la usabilidad y el rendimiento de las aplicaciones al manejar grandes conjuntos de datos. 

- **Paginación** divide el conjunto de datos en páginas más pequeñas, permitiendo al cliente solicitar una página específica de datos.
- **Filtros** permiten al cliente restringir los resultados basándose en ciertos criterios, como un nombre o una fecha específica.

## Archivos Afectados

### 1. `internal/asteroid/handler.go`

Este archivo contiene los controladores HTTP que manejan las solicitudes entrantes relacionadas con los asteroides. Hemos añadido soporte para paginación y filtros en el controlador `GetAllAsteroids`.

#### Cambios Clave

- Se añadieron parámetros de consulta (`page`, `limit` y `name`) para manejar la paginación y los filtros.
- Se implementaron lógica para convertir estos parámetros y usarlos en las consultas a la base de datos.

### 2. `internal/asteroid/repository.go`

Este archivo contiene la lógica de acceso a la base de datos para los asteroides. Hemos añadido un nuevo método `GetAsteroids` que soporta paginación y filtros.

#### Cambios Clave

- Se añadió el método `GetAsteroids` que acepta filtros BSON, límites y saltos (`skip`), y devuelve una lista paginada de asteroides.

### 3. `cmd/main.go`

Este archivo contiene el punto de entrada principal de la aplicación y la configuración de rutas. No se realizaron cambios específicos aquí para la paginación y los filtros, pero es importante verificar que las rutas estén correctamente configuradas.

## Uso de la Paginación y Filtros

"test_pag_filter.md"
