#!/bin/bash

# if arg == up then start dev docker-compose containers
# if arg == down then stop dev docker-compose containers

if [ "$1" == "up" ]; then
docker compose -f docker-compose.dev.yml up -d
elif [ "$1" == "down" ]; then
docker compose -f docker-compose.dev.yml down
else
echo "Operation not supported"
fi
