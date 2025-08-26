# TP2 – Docker: Go + React + PostgreSQL

## 📌 Requisitos
- Docker Desktop o Docker Engine + docker compose
- Cuenta en Docker Hub (imágenes publicadas en `agus12345`)

---

## 🏗️ 1. Construir las imágenes
Ejecutar desde la raíz del proyecto:

```bash
# Backend (Go)
docker build -f Dockerfile.backend -t agus12345/go-api:dev .

# Frontend (React + Nginx)
docker build -f Dockerfile.frontend -t agus12345/react-web:dev .

# Crear versión estable
docker tag agus12345/go-api:dev    agus12345/go-api:v1.0
docker tag agus12345/react-web:dev agus12345/react-web:v1.0


##🚀 2. Ejecutar los contenedores
docker compose up -d
docker compose ps
##Servicios y puertos:
	•	QA → Web: http://localhost:8000 | API: http://localhost:8080
	•	PROD → Web: http://localhost:8001 | API: http://localhost:8081
	•	DB (Postgres): host localhost, puerto 5432



##🌐 3. Acceder a la aplicación
	•	QA: http://localhost:8000
	•	PROD: http://localhost:8001

Endpoints API:
curl http://localhost:8000/api/health   # QA
curl http://localhost:8001/api/health   # PROD


##🗄️ 4. Conectarse a la base de datos

Credenciales (definidas en docker-compose.yml / .env.*):
	•	Host: localhost
	•	Puerto: 5432
	•	DB: appdb
	•	Usuario: appuser
	•	Password: apppass

Conectarse con psql desde el host:
psql "postgresql://appuser:apppass@localhost:5432/appdb"

Desde dentro del contenedor:
docker exec -it $(docker compose ps -q db) psql -U appuser -d appdb


##✅ 5. Verificar funcionamiento

5.1 Salud
curl -i http://localhost:8000/api/health
curl -i http://localhost:8001/api/health


5.2 CRUD mínimo
# Listar (vacío)
curl -i http://localhost:8000/api/items

# Insertar
curl -i -X POST -H "Content-Type: application/json" \
  -d '{"name":"manzana"}' http://localhost:8000/api/items

# Verificar lista
curl -i http://localhost:8000/api/items


5.3 Persistencia
docker compose restart db
sleep 5
curl -i http://localhost:8000/api/items   # los datos siguen presentes



##📦 6. Imágenes en Docker Hub
	•	Backend → agus12345/go-api:{dev,v1.0}
	•	Frontend → agus12345/react-web:{dev,v1.0}