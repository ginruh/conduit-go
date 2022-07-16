#!/bin/bash

if [ "$1" == "init" ]; then
swag init -g cmd/api/main.go internal/controllers/auth.go
elif [ "$1" == "fmt" ]; then
swag fmt -g cmd/api/main.go internal/controllers/auth.go
else
echo "Unsupported operation"
fi