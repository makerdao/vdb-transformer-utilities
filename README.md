# vdb-transformer-utilities

This is a utility library with shared functions that are helpful in writing your own transformers.

## Running tests

Tests can be run with `go run github.com/onsi/ginkgo/ginkgo -r`. 

## Dockerfiles

The docker images are base images used to build and run vdb processes (such as `execute` and `header-sync`). They are base images and are not useful on their own. However the docker-compose file provides a fully running vulcanize setup that you can use as-is or customize for your own needs.
