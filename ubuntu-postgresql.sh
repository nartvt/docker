#!/bin/bash

status=$1

echo $status

docker_compose=docker-compose-postgresql.yml

if [ "$status" = "start" ]; then
    # sudo docker-compose -f $docker_compose --env-file=./environment/.env pull
    # WARNING: Image for service app was built because it did not already exist. 
    # To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
    # sudo docker-compose -f $docker_compose --env-file=./environment/.env up --build -d --remove-orphans
    sudo docker-compose -f $docker_compose --env-file=.env  --project-name=postgresql-ubuntu up -d --remove-orphans
fi

if [ "$status" = "stop" ]; then
    sudo docker-compose -f $docker_compose --env-file=.env down
fi