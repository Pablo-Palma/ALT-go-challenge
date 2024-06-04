### Estructura de un JWT
Un JWT consta de tres partes codificadas en base64 y separadas por puntos (`.`):

1. **Header**: Contiene información sobre el tipo de token y el algoritmo de firma.
2. **Payload**: Contiene las claims, es decir, la información que queremos transmitir (por ejemplo, el nombre de usuario y la fecha de expiración).
3. **Signature**: Es una firma que asegura que el token no ha sido alterado.

Ejemplo de JWT:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

### Proceso de Verificación con `jwt.ParseWithClaims`

#### 1. Parsear el Token
La función `jwt.ParseWithClaims` toma el token en su forma de string y lo divide en sus tres partes: header, payload y signature. Luego decodifica el payload para extraer las claims.

```go
token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
    return jwtKey, nil // jwtKey es la clave secreta usada para firmar el token
})
```

- **tokenStr**: El token JWT recibido.
- **claims**: Una estructura donde se almacenarán las claims extraídas.
- **func(token *jwt.Token) (interface{}, error)**: Una función anónima que devuelve la clave secreta (`jwtKey`) usada para verificar la firma del token.

#### 2. Verificación de la Firma
El algoritmo de firma (especificado en el header del JWT) se usa junto con la clave secreta para generar una firma a partir del header y el payload. Esta firma generada se compara con la firma original que se encuentra en el token.

```go
func(token *jwt.Token) (interface{}, error) {
    return jwtKey, nil // jwtKey es la clave secreta
}
```

- **Header** y **Payload**: Son decodificados y usados para verificar la firma.
- **Signature**: La firma generada se compara con esta parte del token.

Si la firma generada coincide con la firma original, el token es válido y no ha sido alterado.

#### 3. Validación de las Claims
Una vez que se verifica la firma, las claims contenidas en el payload del token son validadas. Esto incluye validaciones automáticas como la expiración del token (`exp`).

```go
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
```

La estructura `Claims` incluye tanto claims personalizadas (como `Username`) como claims estándar (como `ExpiresAt`).

### Ejemplo Completo de Verificación

#### Definición de Claims
```go
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
```

#### Generación del JWT
```go
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
```

#### Verificación del JWT
```go
func VerifyJWT(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil // Clave secreta para verificar la firma
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}
```

#### Middleware de Autorización
```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        claims, err := VerifyJWT(tokenStr)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "username", claims.Username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### Resumen del Proceso

1. **Parseo del Token**: `jwt.ParseWithClaims` decodifica el token en sus tres partes: header, payload y signature.
2. **Verificación de la Firma**: Usa el algoritmo especificado y la clave secreta para generar una firma y la compara con la firma original.
3. **Validación de Claims**: Extrae y valida las claims del payload. Si la firma y las claims son válidas, el token es aceptado.

Este proceso asegura que el token JWT es auténtico y no ha sido modificado, y permite al servidor autorizar al usuario basado en la información contenida en el token.
