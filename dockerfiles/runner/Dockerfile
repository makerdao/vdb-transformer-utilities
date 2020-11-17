# Base image for building the docker containers for building vulcanize projects.
# Meant to be both an optimization of the build, and to remove duplication of the
# various setup steps in other Dockerfiles. 
FROM golang:1.15.5-buster as builder

RUN apt-get update \
        && apt-get install -y \
        ca-certificates \
        bash \
        busybox \
        postgresql-client \
        && update-ca-certificates 2>/dev/null || true
