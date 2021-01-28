# System Docker Compose

This docker-compose file is meant to provide an easy way to setup the vulcanize system without requiring building any custom code. This setup provides:

* A Geth Node with state-diffs emitted
* A header-sync service
* An extract-diffs service with maker storage transformers.
* An execute service with the default maker transforms [mcd-transformers](https://github.com/makerdao/vdb-mcd-transformers).
* An execute service with the oasis transforms [oasis](https://github.com/makerdao/vdb-oasis-transformers).
* A postgraphile instance for API based queries of the data [postgraphile](https://github.com/makerdao/vdb-postgraphile).
* And finally a postgres database.

Running all of this via docker-compose will give you a very similar setup to the public api.


## Dependencies
 - Docker (the rest is handled by docker-compose)

## System Startup

You can start the system by running 

```bash
docker-compose -f dockerfiles/docker-compose.yml up

```

This will work but keep in mind the provided geth-statediffing image will need to sync from the beginning of the chain, and at the time of this writing the provided header-sync will not start updating until block number 71783773. At that point it will begin syncing headers to your database on each block, and will not start syncing storage diffs until some time after that. This means it will need to run for a long time before being viable.

## Customization

The most straightforward way to customize the setup is probably to just edit the file yourself. If you know how to modify docker-compose then you can just get started. 

### Graphile License

If you have a pro license for Graphile, then you can specify it via an environment variable, so:

```bash
GRAPHILE_LICENSE=<license> docker-compose -f dockerfiles/docker-compose.yml up

```

### Adding Different Plugins

If you have written your own plugins for vulcanize you can customize the `docker-compose.yml` file to use them. Keep in mind:

* The images in the default docker-compose file refer to dockerhub. If you build a dockerfile locally you can use it, but you will need to make sure the image is already created on and on your machine.
* Any images you use should be built in the same way as the current plugins. See `https://github.com/makerdao/vdb-mcd-transformers/blob/prod/dockerfiles/execute/Dockerfile` as an example.
* All of the services depend on the `wait-for-it.sh` script which ensures the database is ready to receive connections before attempting to migrate the database.





