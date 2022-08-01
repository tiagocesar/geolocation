# Geolocation service

**Author:** Tiago César Oliveira

This solution is a data importer/API for access to geolocation data.

It comprises two services:

- `importer`, that will look for a dump file on the root of the project with default name `data_dump.csv` and import the valid contents of this file to a postgres database. The service also setups a GRPC interface, so it can interact with other services;
- `api`, that provides a REST API that can be used to consume geolocation data.

## Running the services

The easiest way is to run the `make run` command. It will execute a docker compose file and make all services available. Then the easiest way is to call `http://localhost:8081/locations/{ip}` to interact with the persisted data (not all data will be available on first run, depending on how big the data to be imported is)

Other ways of running the services are available:
- The `importer` can be run via `make run-importer` - it will first guarantee that the postgres container is up before proceeding;
- The `api` can be run via `make run-api`. It will try to interact with the `importer` service only when a request is made;
  - The endpoint `http://localhost:8081/health` is also available (for healthchecks).

## Running tests

- Unit tests can be run via `make unit-tests`;
- Integration tests can be run via `make integration-tests`;
  - Integration tests have special interfaces for data cleanup that aren't available in the general program.

## Points of improvement

Given more time and if this project was a real-world one, some other things would be checked:

- The GRPC server should run over HTTPS;
- Service orchestration isn't really ideal to be run from a compose file to real-world scenarios; deployment scripts aren't defined;
- The services don't implement authentication.