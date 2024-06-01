# Pruebas de Paginación y Filtros con curl

## Ejemplos Prácticos con curl usando Cookies para la Autenticación

### Obtener asteroides con paginación (primera página, 2 asteroides por página)

```sh
curl -X GET "http://localhost:8080/api/v1/asteroides?page=1&limit=2" -b cookies.txt
```

### Obtener asteroides con paginación y filtro por nombre

```sh
curl -X GET "http://localhost:8080/api/v1/asteroides?page=1&limit=2&name=Test" -b cookies.txt
```

### Obtener el segundo lote de asteroides (segunda página, 2 asteroides por página)

```sh
curl -X GET "http://localhost:8080/api/v1/asteroides?page=2&limit=2" -b cookies.txt
```

## Respuestas Esperadas

### Solicitud de Paginación Simple
Debe devolver un array JSON con los asteroides correspondientes a la primera página.

### Solicitud de Paginación con Filtro
Debe devolver un array JSON con los asteroides que coincidan con el nombre filtrado.

### Solicitud del Segundo Lote
Debe devolver un array JSON con los asteroides correspondientes a la segunda página.
