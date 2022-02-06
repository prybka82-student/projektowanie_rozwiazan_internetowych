FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8081

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main