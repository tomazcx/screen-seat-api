# Screen Seat API

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

To list all the available commands, run:

```bash
make help
```

Which will display:

```bash
Usage: make <target>
  help                       Prints available commands
  test                       Run all application's tests
  new-migration              Create new migration
  migrate-up                 Run migrations up
  migrate-down               Run migrations down 1
  up                         Run the application containers and its dependencies
  build                      Run and build the containers
  down                       Stop and destroy the containers
  down-rmi                   Stop and destroy the containers and images
```
