### Logging en Go

El logging es el proceso de registrar mensajes que proporcionan información sobre la ejecución de una aplicación. Estos mensajes pueden incluir información sobre eventos normales, advertencias, errores y otros tipos de información que pueden ser útiles para el desarrollo, depuración y mantenimiento del software.

### ¿Por Qué Usamos Logrus?

Logrus es una biblioteca de logging para Go que ofrece funcionalidades avanzadas en comparación con la biblioteca estándar `log`. Hemos optado por usar Logrus debido a sus múltiples ventajas.

### Beneficios de Usar Logrus

1. **Formatos de Log Avanzados:**
   - Permite la configuración de formatos de log como JSON, lo que es útil para sistemas de análisis de logs y dashboards.

2. **Campos Estructurados:**
   - Soporta campos estructurados, lo que permite añadir contexto adicional a los mensajes de log, facilitando el análisis.

3. **Niveles de Log Granulares:**
   - Ofrece niveles de log más granulares (`Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`), proporcionando un control más preciso sobre el nivel de detalle de los logs.

4. **Hooks:**
   - Permite el uso de hooks para enviar logs a múltiples destinos (por ejemplo, archivos, sistemas de monitoreo, etc.).

5. **Modularidad y Extensibilidad:**
   - Es altamente modular y permite agregar fácilmente funcionalidades adicionales mediante hooks y extensiones.
