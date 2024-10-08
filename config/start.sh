#!/bin/bash

echo "building vendor..."
cd ${WORKDIR} && go mod tidy && go mod vendor

echo "change permissions..."
cd ${WORKDIR} && chown -R developer: .

echo "start server..."
cd ${WORKDIR}/src && go run cmd/api/main.go