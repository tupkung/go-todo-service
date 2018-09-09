FROM golang:1.11-alpine3.8 AS build-env

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git


RUN mkdir -p /go/src/github.com/tupkung/go-todo-service/ && chown tupkung /go/src/github.com/tupkung/go-todo-service/


WORKDIR /go/src/github.com/tupkung/go-todo-service/
ADD . /go/src/github.com/tupkung/go-todo-service

RUN go get ./...

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/src/github.com/tupkung/go-todo-service/todoservice ./cmd/todo/*.go

# final stage
FROM scratch

WORKDIR /

COPY --from=build-env /go/src/github.com/tupkung/go-todo-service/certs/localhost.* /
COPY --from=build-env /go/src/github.com/tupkung/go-todo-service/todoservice /

EXPOSE 8080

CMD ["./todoservice"]

# sudo docker build -f todoservice.Dockerfile -t todo-service:1.0.0 .

# sudo docker run -d --name todo-service -p 8080:8080 -e TODO_SERVICE_ADDR=:8080 -e TODO_SERVICE_CERT_FILE=./localhost.crt -e TODO_SERVICE_KEY_FILE=./localhost.key todo-service:1.0.0