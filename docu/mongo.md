#### Configuración de la Base de Datos y Colección

1. **Acceder a MongoDB Shell**:

   Inicia MongoDB Shell para conectarte a la base de datos MongoDB en ejecución:

   ```sh
   mongosh
   ```

2. **Crear y Seleccionar la Base de Datos**:

   Una vez dentro de `mongosh`, crea y selecciona la base de datos `asteroidsdb`:

   ```sh
   use asteroidsdb
   ```

3. **Crear una Colección**:

   Crea una colección llamada `asteroids` dentro de `asteroidsdb`:

   ```sh
   db.createCollection("asteroids")
   ```

#### Comprobación de la Configuración

Para verificar que la base de datos y la colección se han creado correctamente:

1. **Mostrar las Bases de Datos**:

   ```sh
   show dbs
   ```

2. **Mostrar las Colecciones en `asteroidsdb`**:

   ```sh
   show collections
   ```

#### Ejecución de MongoDB con Docker

Para ejecutar MongoDB en un contenedor Docker:

1. **Ejecutar MongoDB en Docker**:

   Utiliza el siguiente comando para iniciar MongoDB en un contenedor Docker:

   ```sh
   docker run -d --name mongodb -p 27017:27017 mongo
   ```

   - **`docker run`**: Comando para crear y ejecutar un contenedor.
   - **`-d`**: Ejecuta el contenedor en modo desacoplado (detached), es decir, en segundo plano.
   - **`--name mongodb`**: Asigna un nombre al contenedor (`mongodb`).
   - **`-p 27017:27017`**: Mapea el puerto 27017 del contenedor al puerto 27017 del host, permitiendo el acceso a MongoDB.
   - **`mongo`**: Especifica la imagen de Docker a utilizar (en este caso, la imagen oficial de MongoDB).

2. **Acceder a MongoDB en Docker**:

   Puedes acceder a MongoDB en el contenedor Docker usando `mongosh`:

   ```sh
   mongosh "mongodb://localhost:27017"
   ```

3. **Seleccionar la Base de Datos `asteroidsdb`**:

   ```sh
   use asteroidsdb
   ```

### Resumen de Comandos

1. **Instalar MongoDB y mongosh**:

   ```sh
   brew tap mongodb/brew
   brew install mongodb-community
   brew install mongosh
   ```

2. **Iniciar MongoDB**:

   ```sh
   brew services start mongodb/brew/mongodb-community
   ```

3. **Acceder a MongoDB Shell y Configurar la Base de Datos**:

   ```sh
   mongosh
   use asteroidsdb
   db.createCollection("asteroids")
   ```

4. **Ejecución de MongoDB en Docker**:

   ```sh
   docker run -d --name mongodb -p 27017:27017 mongo
   mongosh "mongodb://localhost:27017"
   use asteroidsdb
   ```

### Comandos para Verificar la Base de Datos y Colecciones

1. **Acceder a MongoDB en el contenedor Docker**:

   ```sh
   docker exec -it mongodb mongosh
   ```

2. **Mostrar todas las bases de datos**:

   ```sh
   show dbs
   ```

3. **Seleccionar la base de datos `asteroidsdb`**:

   ```sh
   use asteroidsdb
   ```

4. **Mostrar todas las colecciones en `asteroidsdb`**:

   ```sh
   show collections
   ```

5. **Listar documentos en la colección `asteroids`**:

   ```sh
   db.asteroids.find().pretty()
   ```

6. **Listar documentos en la colección `users`**:

   ```sh
   db.users.find().pretty()
   ```
