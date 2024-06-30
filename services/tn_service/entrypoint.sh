#!/bin/sh

alembic upgrade head
if [ $? -eq 0 ]; then
    echo "Migrations - OK"
    python main.py
else
    echo "Migrations failed"
    exit 1
fi

exec "$@"