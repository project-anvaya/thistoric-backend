#!/usr/bin/env bash
set -e
curl -s http://localhost:8080/health | jq .
curl -s "http://localhost:8080/vendors?category=Caterers&location=Pune" | jq .
curl -s http://localhost:8080/vendors/1 | jq .
