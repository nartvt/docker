# access to console by docker command 
    docker exec -it <container id or container name> /bin/bash

# access with root user
    docker exec -u 0 -it <container id or container name> /bin/bash

# stop all container
    docker stop $(docker ps -a -q)

# remove all of container 
    docker rm $(docker ps -a -q)

# Alias ll in the case just only use ls cannot use ll
    alias ll="ls -la"

# Global install composer(https://getcomposer.org/doc/00-intro.md)
    1. go to : https://getcomposer.org/download/ , download compose.phar
    2. execute command: mv composer.phar /usr/local/bin/composer
    
 access to docker container execute, database/app:  composer install

composer install
php artisan migrate --seed
php artisan migrate

rabbitmqctl add_user justrace p4ssw0rd
rabbitmqctl set_user_tags justrace administrator