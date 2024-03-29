FROM golang:1.19.5-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o todo-api ./cmd/main.go

EXPOSE 8000

CMD ["./todo-api"]