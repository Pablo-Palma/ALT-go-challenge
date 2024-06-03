# Dashboard de Métricas de la API de Registro de Asteroides

## Introducción a Grafana

**Grafana** es una plataforma de código abierto para la visualización y monitoreo de métricas, que permite a los usuarios consultar, visualizar y alertar sobre datos de series temporales. Es ampliamente utilizada para observar la infraestructura y aplicaciones de TI. Grafana soporta múltiples fuentes de datos, incluyendo Prometheus, InfluxDB, Graphite, Elasticsearch, MySQL, y muchos más.

### Beneficios de Grafana

- **Visualización**: Proporciona una interfaz gráfica rica para crear dashboards personalizados que muestran métricas y datos en tiempo real.
- **Alertas**: Permite configurar alertas basadas en las métricas monitoreadas para notificar sobre eventos críticos.
- **Integración**: Se integra con múltiples fuentes de datos, lo que facilita la recopilación y visualización de métricas de diferentes sistemas.
- **Facilidad de uso**: Ofrece una interfaz intuitiva que facilita la creación y gestión de dashboards sin necesidad de conocimientos avanzados en programación.

### Integración en el Proyecto

La integración de Grafana en el proyecto se refleja en varios archivos de configuración y documentación:

1. **Archivo `docker-compose.yml`**: Define los servicios de Docker, incluyendo Grafana, Prometheus y la aplicación Go.
2. **Archivo de Configuración de Prometheus (`prometheus.yml`)**: Configura los targets que Prometheus debe monitorear.
3. **Código de la Aplicación (`main.go`)**: Implementa las métricas de Prometheus en la aplicación Go.
4. **Documentación (`grafana.md`)**: Detalla la configuración y uso de dashboards en Grafana.

### Configuración de Grafana

#### Acceder a Grafana

Abre Grafana en tu navegador web en `http://localhost:3000` y accede con las credenciales configuradas (por defecto, admin/admin).

#### Configurar Prometheus como Fuente de Datos

1. Ve a "Configuration" > "Data Sources".
2. Añade una nueva fuente de datos y selecciona "Prometheus".
3. Configura la URL de Prometheus (por ejemplo, `http://prometheus:9090`) y guarda la configuración.

## Panel de Métricas HTTP

**Nombre del Panel:** Métricas HTTP

**Descripción del Panel:** Este panel muestra las métricas de solicitudes HTTP para la API de registro de asteroides, incluyendo el total de solicitudes por endpoint, la tasa de solicitudes por segundo y la distribución de solicitudes por método.

**Consultas y Tipos de Visualización:**

- **Total de Solicitudes HTTP por Endpoint:**
  - **Query:**
    ```promql
	sum by (endpoint) (rate(http_requests_total[1m]))
    ```
  - **Tipo de Visualización:** Bar chart

- **Tasa de Solicitudes HTTP por Segundo para un Endpoint Específico:**
  - **Query:**
    ```promql
    rate(http_requests_total{endpoint="/api/v1/asteroides"}[1m])
    ```
  - **Tipo de Visualización:** Line chart

- **Número Total de Solicitudes HTTP Agrupadas por Método:**
  - **Query:**
    ```promql
	sum by (method) (rate(http_requests_total[1m]))
    ```
  - **Tipo de Visualización:** Bar chart

## Panel de Métricas de Rendimiento

**Nombre del Panel:** Métricas de Rendimiento

**Descripción del Panel:** Este panel muestra las métricas de rendimiento para la API de registro de asteroides, incluyendo el uso de memoria, el número de goroutines en ejecución y los tiempos de respuesta HTTP.

**Consultas y Tipos de Visualización:**

- **Uso de Memoria:**
  - **Query:**
    ```promql
    go_memstats_alloc_bytes
    ```
  - **Tipo de Visualización:** Line chart

- **Número de Goroutines:**
  - **Query:**
    ```promql
    go_goroutines
    ```
  - **Tipo de Visualización:** Line chart

- **Tiempos de Respuesta HTTP (si tienes esta métrica implementada):**
  - **Query:**
    ```promql
    histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le))
    ```
  - **Tipo de Visualización:** Line chart

