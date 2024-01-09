FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
