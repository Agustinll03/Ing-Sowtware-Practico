# Decisiones TP2 – Go + React + PostgreSQL

## 1. Elección de la app
- **Qué hicimos**: API en Go (items) + frontend React (Vite) + DB PostgreSQL.
- **Por qué**: muestra FE/BE/DB, es liviano y fácil de contenerizar.
- **Alternativa**: Node/Express + React. No se eligió porque la imagen final sería más pesada.

## 2. Dockerfiles
- **Backend**: multi-stage (`golang:1.22-alpine` → `alpine:3.20`), binario estático, usuario no-root.
- **Frontend**: build con `node:20-slim`, runtime con `nginx:stable-alpine`.
- **Por qué**: multi-stage reduce tamaño; Nginx es estándar para estáticos.
- **Alternativas**: `scratch` o `distroless` → más chicos pero difíciles de depurar.

## 3. Base de datos
- PostgreSQL 16 con volumen `db_data` para persistencia.
- **Por qué**: DB real en red, healthcheck con `pg_isready`.
- **Alternativa**: SQLite (más simple) pero no demuestra conexión entre contenedores.

## 4. QA y PROD con misma imagen
- Config con `.env.qa` y `.env.prod` para el backend.
- Nginx monta `default.qa.conf` o `default.prod.conf`.
- **Por qué**: una sola imagen, cambia solo configuración.
- **Alternativa**: imágenes distintas para QA/PROD, pero rompe reproducibilidad.

## 5. docker-compose
- Servicios: `db`, `api-qa`, `api-prod`, `web-qa`, `web-prod`.
- Mapeo de puertos: QA en 8000/8080, PROD en 8001/8081.
- **Por qué**: correr QA y PROD a la vez demuestra misma imagen/diferente config.

## 6. Versionado y publicación
- Tags en Docker Hub: `:dev` (iteración) y `:v1.0` (estable).
- **Por qué**: reproducibilidad del entorno final.
- **Alternativa**: SemVer completo (`v1.0.1`, `v1.1.0`) → más fino, pero suficiente con `v1.0`.

## 7. Problemas y fixes principales
- **React is not defined** → agregar `import React`.
- **/api/health 404** → corregir `proxy_pass` en Nginx (sin `/`).
- **GET /api/items = null** → inicializar slice vacío en Go.
- **Compose error** → mover servicios bajo `services:`.
- **Push denegado** → crear repos en Docker Hub, login con token y retag al usuario correcto.

## 8. Calidad y trazabilidad
- Commits atómicos + PRs.
- Test reproducible (`test.sh`).
- En un equipo real: CI/CD que buildée imágenes, corra tests y publique con tags.
