# Base image for building the docker containers for building vulcanize projects.
# Meant to be both an optimization of the build, and to remove duplication of the
# various setup steps in other Dockerfiles.
FROM golang:1.17.3-bullseye as builder

RUN apt-get update && \
    apt-get install -y \
    make git g++ linux-headers-amd64

ENV GO111MODULE on

RUN go get -tags='no_mysql no_sqlite3 no_mssql no_redshift' github.com/pressly/goose/cmd/goose@v2.7.0-rc5

