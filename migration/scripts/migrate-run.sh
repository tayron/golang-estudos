#!/bin/bash

migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${DATABASE_NAME}" up