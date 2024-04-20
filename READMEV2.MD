# Backend Engineering Interview Assignment (Golang)

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.19
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

    Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project

To start working, execute

```
make go-module 
```

To generate swagger api definitions

```
make swag
```

To generate mocks need install mockery on your platform

```
make mock
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

You should be able to access the API Definitions at http://localhost:8080/swagger/index.html

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

To run test, run the following command:

```
make test
```
