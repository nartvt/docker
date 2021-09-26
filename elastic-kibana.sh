#!/bin/bash
# docker pull elasticsearch:7.14.1

# docker network create lozi_network

# docker run -d --name elasticsearch --net lozi_network -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.14.1

# docker run -d --name kibana --net lozi_network -p 5601:5601 -e "ELASTICSEARCH_HOSTS=http://elasticsearch:9200" kibana:7.14.1


docker_compose=docker-compose-lozi.yml
docker-compose -f $docker_compose  --project-name=lozi-docker up -d --remove-orphans