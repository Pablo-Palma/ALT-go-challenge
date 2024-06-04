## Headers HTTP

### ¿Qué es un Encabezado HTTP?
Un encabezado HTTP es una parte de la solicitud o respuesta HTTP que proporciona metadatos adicionales sobre la transacción HTTP. Los encabezados son pares clave-valor enviados en el formato:
```
Clave: Valor
```

### Tipos Comunes de Encabezados HTTP

#### Encabezados de Solicitud (Request Headers)
- **Authorization**: Incluye credenciales para autenticar al cliente con el servidor.
- **Content-Type**: Indica el tipo de contenido que se está enviando al servidor (e.g., `application/json`).
- **Accept**: Indica qué tipo de respuesta espera el cliente (e.g., `application/json`).

#### Encabezados de Respuesta (Response Headers)
- **Content-Type**: Indica el tipo de contenido de la respuesta del servidor.
- **Set-Cookie**: Establece cookies que el navegador almacenará y enviará en futuras solicitudes.

### Encabezado de Autorización

#### ¿Qué es un Token en el Encabezado de Autorización?
Un token es una cadena de caracteres utilizada para identificar a un usuario autenticado. Los tokens se pueden enviar en el encabezado `Authorization` para autenticar solicitudes HTTP. Un esquema común es el Bearer, utilizado para tokens JWT (JSON Web Tokens).

#### Formato del Encabezado de Autorización
El token se envía en el encabezado `Authorization` de la solicitud HTTP utilizando el esquema Bearer:
```
Authorization: Bearer <token_jwt>
```

### Uso de Tokens en el Desarrollo Web
- **Autenticación de API**: Los tokens se utilizan para autenticar solicitudes a APIs RESTful. Al incluir el token en el encabezado `Authorization`, el servidor puede verificar la identidad del cliente y autorizar la solicitud.
- **Seguridad**: Los tokens son más seguros que las cookies porque no se almacenan en el navegador. Además, los tokens pueden tener un tiempo de expiración, limitando su validez temporal.
- **Estándar en el Desarrollo Web**: Usar tokens en el encabezado `Authorization` es una práctica estándar en el desarrollo web moderno, especialmente para APIs.

