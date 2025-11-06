#!/usr/bin/env bash
set -euo pipefail
export $(grep -v '^#' .env | xargs)
go run ./cmd/api
