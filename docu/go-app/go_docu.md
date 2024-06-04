### Go Tools

**Context**(context.Context), es una herramienta para manejar datos de solicitud, señales de cancelación, y timeouts en aplicaciones concurrentes.
Se utiliza comunmente en aplicaiones web o de red para pasar información entre diferentes manejadores y middlewares de una solicitud.

### Go Sintaxt

**asignación múltiple en Go:** hashedPassword, err := HashPassword(user.Password), la funcion `HashPassword` devuleve dos valores: el hash de la password y el posble error a la variable err.

### Paquetes en Go

En Go, los paquetes agrupan código relacionado y permiten su reutilización. Declaramos el paquete en la parte superior del archivo (`package asteroid`). Cuando queramos usar la estructura `Asteroid` en otros archivos, importaremos este paquete, similar a cómo se hace en C con las headers.

Por ejemplo, si queremos usar `Asteroid` en otro archivo, añadimos:
```go
import "e42-go-challenge/internal/asteroid"
```

Luego, podemos crear un nuevo asteroide:
```go
newAsteroid := asteroid.Asteroid{
    ID: "123",
    Name: "Ceres",
    Diameter: 939.4,
    DiscoveryDate: "1801-01-01",
    ObservationCount: 100,
    DistanceFromEarth: 2.77,
}
```
