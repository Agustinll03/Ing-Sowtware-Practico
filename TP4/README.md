# TP4 – Azure DevOps Pipelines (Self-Hosted)

Monorepo con **front (React + Vite)** y **back (Go 1.22)**. Pipeline CI corre en **agente self-hosted**.

## Estructura
- `front/`: React + Vite (`npm ci && npm run build`)
- `back/`: Go 1.22 (`go build -o out/app .`)
- `azure-pipelines.yml`: pipeline multi-job que publica artefactos `front-dist` y `back-bin`.

## Correr local
```bash
# Front
cd front
npm ci
npm run dev  # http://localhost:5173

# Back
cd back
go run main.go  # http://localhost:8080/health