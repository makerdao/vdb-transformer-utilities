#! /usr/bin/env bash

set -e

function message() {
    echo
    echo -----------------------------------
    echo "$@"
    echo -----------------------------------
    echo
}

TAG=latest
BUILDER_NAME=vdb-builder
RUNNER_NAME=vdb-runner

message BUILDING BUILDER DOCKER IMAGE
docker build -f dockerfiles/builder/Dockerfile . -t makerdao/$BUILDER_NAME:$TAG

message BUILDING RUNNER DOCKER IMAGE
docker build -f dockerfiles/runner/Dockerfile . -t makerdao/$RUNNER_NAME:$TAG

message LOGGING INTO DOCKERHUB
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USER" --password-stdin

message PUSHING BUILDER DOCKER IMAGE
docker push makerdao/$BUILDER_NAME:$TAG

message PUSHING RUNNER DOCKER IMAGE
docker push makerdao/$RUNNER_NAME:$TAG
