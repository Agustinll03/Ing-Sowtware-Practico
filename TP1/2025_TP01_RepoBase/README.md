# Proyecto Base para Práctica de Git

Este proyecto es parte del Trabajo Práctico 01 – Git Básico.

Contiene archivos simples para que puedas practicar creación de ramas, commits, solución de errores y versionado.

## Estructura
- `src/app.js`: contiene un script básico.
- `data/info.txt`: contiene datos de ejemplo.

## Uso rápido

## Para levantar todos los servicios en segundo plano:
docker compose up -d

## Para ver el estado de los contenedores:
docker compose ps

## Para ver logs de un servicio (ejemplos):
docker compose logs -f db
docker compose logs -f api-qa
docker compose logs -f api-prod
docker compose logs -f web-qa
docker compose logs -f web-prod

## Para acceder a la aplicación:

	•	Web QA → http://localhost:8000
	•	Web PROD → http://localhost:8001
	•	API QA → http://localhost:8080
	•	API PROD → http://localhost:8081
	•	Base de datos → psql -h localhost -U appuser -d appdb -p 5432

## Para apagar todo (contenedores, pero conserva datos):
docker compose down

## Para apagar todo y borrar también los datos de la base:
docker compose down -v

## Para reconstruir imágenes (si cambiaste Dockerfiles o configuraciones):
docker compose build
