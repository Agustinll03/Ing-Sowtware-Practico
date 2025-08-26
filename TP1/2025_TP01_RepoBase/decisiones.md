# Decisiones TP1 – Flujo de trabajo con Git

## 1) Preparación del repo
- Partí del repo base `ingsoft3ucc/2025_TP01_RepoBase` y configuré mi identidad con `user.name` y `user.email`.
- Agregué remoto `upstream` para poder referenciar el origen del trabajo.

**Alternativa:** trabajar directo en `main`.  
**Por qué NO:** mezcla desarrollo con producción, pierde trazabilidad.

## 2) Desarrollo de funcionalidad
- Rama: `feat/saludo-personalizado`.
- Commits atómicos:
  1. `feat(saludo): prefijo configurable via env SALUDO_PREFIX`
  2. `docs(readme): documentar uso de SALUDO_PREFIX`

**Estrategia:** separar la implementación del cambio de la documentación para poder revertir/iterar de forma independiente.

**Alternativas:** un solo commit grande.  
**Por qué NO:** cuesta revisar, revertir y auditar.

## 3) Corrección de error (hotfix)
- Simulé un bug en `main` (saludo mal escrito).
- Creé `hotfix/arreglar-saludo`, hice el fix, lo integré a `main` con merge sin fast-forward y luego a la rama de desarrollo con `cherry-pick`.

**Por qué:** el hotfix debe llegar primero a producción y después sincronizarse con la rama de trabajo sin mezclar otros cambios.

**Alternativas:**
- Merge de `main` → `feat/*`.  
  **Contras:** trae también commits ajenos al fix.
- Rehacer el fix a mano en `feat/*`.  
  **Contras:** duplica trabajo y arriesga inconsistencias.

## 4) Pull Request
- Abrí un PR `feat/saludo-personalizado → main` y lo fusioné.  
- Justificación: revisión, trazabilidad y registro formal del cambio.

**Alternativa:** push directo a `main`.  
**Por qué NO:** sin revisión ni historial de decisión.

## 5) Versión etiquetada
- Tag: `v1.0` en `main` después de integrar la feature y el hotfix.  
- Convención: SemVer simplificado (mayor.menor.parche), empezando por `v1.0` (primera release estable).

**Alternativa:** fechas en tags (ej. `2025-08-26`).  
**Por qué NO:** menos expresivo respecto a compatibilidad.

## 6) Trazabilidad y calidad (cómo lo haría en equipo)
- Ramas por feature/hotfix, commits atómicos con convención de mensajes.
- PRs obligatorios, CI con tests/lint y reglas de protección sobre `main`.
- Tags por release y changelog.

