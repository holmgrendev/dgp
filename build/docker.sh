#!/bin/bash

start(){
    echo "_____ COPYING FILES _____" >&2

    # Create a temporary build directory
    cd $(dirname "$0")/docker
    mkdir temp

    # Go to project root and copy directories and files
    cd ../../
    cp -a cmd internal public go.mod ./build/docker/temp

    # Go to docker build directory and copy files
    cd ./build/docker
    cp .env docker-compose.yml Dockerfile ./temp

    # Go to temp build directory and build project
    cd temp

    echo "_____ COMPOSING CONTAINERS _____" >&2
    sudo docker compose up --build --force-recreate --no-deps --detach

    # Go back and remove temp build directory
    cd ../
    rm -rf temp
    cd ../
}

kill(){
    echo "_____ STOPPING CONTAINERS _____" >&2
    sudo docker stop dgp-web
    sudo docker stop dgp-db

    echo "_____ REMOVING CONTAINERS _____" >&2
    sudo docker rm dgp-web
    sudo docker rm dgp-db
}

while getopts 'skr' OPTION; do
  case "$OPTION" in
    s)
      start
      exit 1
      ;;
    k)
      kill
      exit 1
      ;;
    r)
      echo "_____ RESTARTING CONTAINERS _____" >&2
      kill
      start
      exit 1
      ;;
  esac
done
echo "Script usage:\n\t-s Build and [s]tart containers\n\t-k [k]ill and remove containers\n\t-r [r]estart containers (Kill, remove, build and start)" >&2
exit 1