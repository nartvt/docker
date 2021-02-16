#!/bin/bash

status=$1

echo $status

if [ "$status" = "up" ]; then
    sudo docker-compose -f docker-compose_app.yml pull
    sudo docker-compose -f docker-compose_app.yml up -d --remove-orphans
fi

if [ "$status" = "down" ]; then
    sudo docker-compose -f docker-compose_app.yml down
fi