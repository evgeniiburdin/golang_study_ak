#!/bin/bash

# Container name
CONTAINER_NAME="postgres"
NETWORK_NAME="backend-network"

if ! docker network ls | grep -q $NETWORK_NAME; then
    echo "Creating network $NETWORK_NAME..."
    docker network create $NETWORK_NAME
fi

# Check if a container with the specified name already exists
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "Container with name $CONTAINER_NAME already exists. Stopping and removing it..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# Start a new container with PostgreSQL
echo "Starting a new container $CONTAINER_NAME..."
docker run --name $CONTAINER_NAME --network=$NETWORK_NAME -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

# Wait for PostgreSQL to initialize
echo "Waiting for PostgreSQL to initialize..."
sleep 5

# Create two databases and tables in them
echo "Creating databases and tables..."
docker exec -i $CONTAINER_NAME psql -U postgres <<EOF
CREATE DATABASE user_db;
\c user_db
CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    user_role VARCHAR(255)
);

CREATE DATABASE auth_db;
\c auth_db
CREATE TABLE IF NOT EXISTS auth (
    email VARCHAR(255) UNIQUE,
    refreshToken VARCHAR(255) UNIQUE
);
EOF

echo "Databases and tables created successfully!"
