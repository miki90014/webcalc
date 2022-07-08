FROM golang:latest 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download -x
COPY *.go ./

RUN go run main.go

EXPOSE 8080
EXPOSE 8081
