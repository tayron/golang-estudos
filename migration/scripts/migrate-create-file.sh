#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <migration-name>"
  exit 1
fi

migrate create -ext sql -dir migrations -seq "$1"