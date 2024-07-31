#!/bin/sh

echo "Building and running store-api-gateway..."
docker build -t store-api-gateway_image ./store-api-gateway
docker run -d --name store-api-gateway_container store-api-gateway_image

echo "Building and running store-orders-service..."
docker build -t store-orders-service_image ./store-orders-service
docker run -d --name store-orders-service_container store-orders-service_image

echo "Building and running store-payments-service..."
docker build -t store-payments-service_image ./store-payments-service
docker run -d --name store-payments-service_container store-payments-service_image

echo "Building and running store-products-service..."
docker build -t store-products-service_image ./store-products-service
docker run -d --name store-products-service_container store-products-service_image

echo "Building and running store-users-service..."
docker build -t store-users-service_image ./store-users-service
docker run -d --name store-users-service_container store-users-service_image

wait
