#!/bin/bash
set -e

echo "==============================="
echo "🔎 Test de TP2 - QA y PROD"
echo "==============================="

echo
echo "▶️  QA - Health"
curl -s -i http://localhost:8000/api/health || true

echo
echo "▶️  QA - Items (inicio)"
curl -s -i http://localhost:8000/api/items || true

echo
echo "▶️  QA - POST /api/items (agrego 'manzana')"
curl -s -i -X POST -H "Content-Type: application/json" \
  -d '{"name":"manzana"}' http://localhost:8000/api/items || true

echo
echo "▶️  QA - Items (después del insert)"
curl -s -i http://localhost:8000/api/items || true

echo
echo "▶️  PROD - Health"
curl -s -i http://localhost:8001/api/health || true

echo
echo "▶️  PROD - Items"
curl -s -i http://localhost:8001/api/items || true

echo
echo "▶️  Persistencia: reinicio DB y vuelvo a consultar"
docker compose restart db
sleep 5
curl -s -i http://localhost:8000/api/items || true

echo
echo "✅ Pruebas finalizadas."