#!/bin/bash

ENGINE=""

# Try to use PODMAN
if [[ "which: no" =~ $(which podman) ]]; then
    echo -e "Podman not found...\nLoking for docker instead."
else
    ENGINE="podman"
fi

# If PODMAN was not found, try to use DOCKER
if [[ $ENGINE = "" ]]; then

    if [[ "which: no" =~ $(which docker) ]]; then
        echo "Docker not found..."
    else
        ENGINE="docker"
    fi

fi

# Quit if neither PODMAN nor DOCKER was found
if [[ $ENGINE = "" ]]; then
    echo -e "No compatible container found in the PATH.\nQuitting..."
    return 1
else
    echo -e "Using $ENGINE".
fi

COMMAND="$ENGINE images"
container_found="false"

# Look for existing container
if [[ "go-borg" =~ $( $COMMAND) ]]; then
    echo "Container found"
    $container_found="true"
else
    echo -e "No container found."
fi

if [[ $container_found = "false" ]]; then
    read -p "Do you want to build it? [y/n]? " ANSWER

    if [[ $ANSWER = "y" ]] || [[ $ANSWER = "Y" ]]; then
        COMMAND="$ENGINE build -t go-borg ."
        $COMMAND
    else
        echo "Quiting.."
        exit
    fi
fi

read -p "Do you want to run the app? [y/n]? " ANSWER
if [[ $ANSWER = "y" ]] || [[ $ANSWER = "Y" ]]; then
    COMMAND="$ENGINE run -p 5000:5000 go-borg"
    $COMMAND
else
    echo "Quiting.."
    exit
fi