# Decisiones TP3 – Azure DevOps

## 1. Metodología
- Elegimos Agile
- Alternativa: Scrum (más burocrático)

## 2. Work Items
- Epic: Plataforma alquileres
- Stories: Login, Publicación, Reserva
- Tasks: Backend y Frontend por Story
- Bugs: validación email, reservas duplicadas
- Sprint 1: 2 semanas

## 3. Repos y Policies
- Repo en Azure Repos
- Ramas feature/*
- PR obligatorio
- Policies: 1 reviewer + build validation

## 4. Pipeline
- Backend Go + Frontend React
- Validación en PRs

## 5. Evidencias
- Boards con Stories y Tasks
- Sprint planificado
- PRs mergeados
- Pipelines exitosos

## 6. Problemas
- Permisos pipeline → se habilitó en Project Settings
- Áreas mal configuradas → se corrigieron
- Policies laxas → se endurecieron

## 7. Calidad y trazabilidad
- Commits atómicos
- PRs revisados
- CI obligatorio
- Boards linkados a PRs
