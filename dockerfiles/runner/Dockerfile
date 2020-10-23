# Base image for building the docker containers for building vulcanize projects.
# Meant to be both an optimization of the build, and to remove duplication of the
# various setup steps in other Dockerfiles. 
FROM golang:1.15.2-buster as builder

RUN apt-get update \
        && apt-get upgrade -y \
        && apt-get install -y \
        ca-certificates \
        bash \
        busybox \
        postgresql-client \
        && update-ca-certificates 2>/dev/null || true

WORKDIR /go/src/github.com/makerdao/vulcanizedb

# Direct logs to stdout for docker log driver
RUN ln -sf /dev/stdout /go/src/github.com/makerdao/vulcanizedb/vulcanizedb.log
