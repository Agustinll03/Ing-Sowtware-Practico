# Decisiones TP1 – Git Básico

## 1. Flujo de trabajo usado
Utilicé un flujo basado en **Git Flow simplificado**:
- `main`: rama estable de producción.
- `feat/*`: ramas de funcionalidad (en mi caso `feat/saludo-personalizado`).
- `hotfix/*`: ramas de corrección urgente (`hotfix/arreglar-saludo`).

**Por qué**: separa el desarrollo de nuevas features del código estable y permite aplicar fixes críticos directamente en `main`.

**Alternativa**: trabajar siempre en `main`.  
**Por qué no**: perdería trazabilidad, mezclaría cambios en desarrollo con producción y dificultaría la revisión.

---

## 2. Integración del fix
El fix se aplicó en una rama `hotfix/arreglar-saludo`.  
Pasos:
1. Detectar bug en `main` y crear la rama hotfix.  
2. Corregir el error y commitear.  
3. Integrar a `main` con un **merge sin fast-forward** para que quede registro explícito del hotfix.  
4. Aplicar también a la rama de desarrollo (`feat/saludo-personalizado`) mediante **cherry-pick**, para mantener ambas ramas coherentes.

**Por qué esta estrategia**:  
- El merge deja historial claro en `main`.  
- El cherry-pick evita traer commits no deseados de `main` a la rama feature.

---

## 3. Problemas y soluciones
- **Simulación de bug en `main`**: al modificar manualmente el código, me aseguré de revertirlo en el hotfix.  
- **Confusión con integración**: evalué usar `merge main → feature`, pero habría traído commits innecesarios; lo resolví con `cherry-pick` del fix.  
- **Posibles conflictos**: no aparecieron, pero si ocurren la estrategia sería resolverlos manualmente y documentarlos en el commit.

---

## 4. Calidad y trazabilidad en un equipo real
- **Commits atómicos y claros**, siguiendo convención tipo Conventional Commits (`feat:`, `fix:`, `docs:`).  
- **Pull Requests obligatorios** para revisión por pares antes de fusionar a `main`.  
- **Tags semánticos** para versiones (`v1.0`, `v1.1.0`).  
- **CI/CD** automatizado: ejecutar tests en cada PR, verificar builds y publicar releases automáticamente.  
- **Ramas protegidas**: evitar pushes directos a `main`, solo merges aprobados.  

Esto asegura que cada cambio sea trazable, auditable y verificable antes de llegar a producción.