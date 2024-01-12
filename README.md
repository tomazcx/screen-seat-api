

A RESTful API made for Screen Seat, a system designed to search for available movies and purchase tickets to watch them.

## STACK

- Golang
- Fiber (Go HTTP framework)
- PostgreSQL
- SQLC
- RabbitMQ
- Docker
- Testify (Testing)
- AWS S3 

## GUIDE TO RUN

To run this project locally, starting by clone the repository:

```bash
git clone https://github.com/tomazcx/screen-seat-api.git
```

Then, cd into the project folder:

```bash
cd screen-seat-api
```

Generate the code for the database queries with [SQLC](https://sqlc.dev/) (make sure to have it installed)

```bash
sqlc generate
```

Finally, start the project by running:

```bash
make up
```

## USAGE

To list all the available `make` commands, run:

```bash
make help
```

## TESTING

You can execute one of the two commands to run all the application tests:

```bash
make test
# or 
go test ./...
```
