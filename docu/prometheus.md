# Configuración de Prometheus para la API de Registro de Asteroides

## Introducción a Prometheus

**Prometheus** es una herramienta de monitoreo y alerta de código abierto que recolecta y almacena métricas como datos de series temporales. Es ampliamente utilizada para monitorear aplicaciones y sistemas.

### Beneficios de Prometheus

- **Recolecta métricas de series temporales.**
- **Permite consultas avanzadas con PromQL.**
- **Soporta alertas basadas en métricas.**
- **Funciona de manera independiente y se integra con múltiples sistemas.**

## Configuración de Prometheus

### Acceder a Prometheus

Abre Prometheus en tu navegador web en `http://localhost:9090`.

### Configuración de Prometheus (`prometheus.yml`)

El archivo de configuración de Prometheus define los targets que debe monitorear. Aquí está el ejemplo actual:

```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
    - targets: ['localhost:9090']

  - job_name: 'go-app'
    static_configs:
    - targets: ['go-app:8080']
```

### Consultas Básicas en Prometheus

#### Total de Solicitudes HTTP por Endpoint

```promql
sum by (endpoint) (http_requests_total)
```

#### Tasa de Solicitudes HTTP por Segundo para un Endpoint Específico

```promql
rate(http_requests_total{endpoint="/api/v1/asteroides"}[1m])
```

#### Uso de Memoria

```promql
go_memstats_alloc_bytes
```

#### Número de Goroutines

```promql
go_goroutines
```

## Integración en el Proyecto

La integración de Prometheus en el proyecto se refleja en varios archivos:

1. **`docker-compose.yml`**: Define los servicios de Docker, incluyendo Prometheus, Grafana y la aplicación Go.
2. **`prometheus.yml`**: Configura los targets que Prometheus debe monitorear.
3. **`main.go`**: Implementa las métricas de Prometheus en la aplicación Go.

### Configuración del Servicio en `docker-compose.yml`

```yaml
prometheus:
  image: prom/prometheus
  container_name: prometheus
  ports:
    - "9090:9090"
  volumes:
    - ./prometheus:/etc/prometheus
  command:
    - '--config.file=/etc/prometheus/prometheus.yml'
```

### Implementación de Métricas en `main.go`

Ejemplo de cómo se exponen las métricas en la aplicación Go:

```go
import (
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "net/http"
)

var httpRequestsTotal = prometheus.NewCounterVec(
  prometheus.CounterOpts{
    Name: "http_requests_total",
    Help: "Total number of HTTP requests",
  },
  []string{"method", "endpoint"},
)

func init() {
  prometheus.MustRegister(httpRequestsTotal)
}

func prometheusMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
    next.ServeHTTP(w, r)
  })
}

func main() {
  http.Handle("/metrics", promhttp.Handler())
  http.ListenAndServe(":8080", nil)
}
```
