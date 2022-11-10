#!/bin/sh

set -e
echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE"

echo "start the app"
exec "$@"
