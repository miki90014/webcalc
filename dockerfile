FROM golang:latest AS builder
ENV GOPROXY="https://repo.cci.nokia.net/proxy-golang-org"
ENV GOSUMDB=off
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o repo -ldflags "-w -extldflags '-static'" .

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/repo .
USER 8000:8000
EXPOSE 8000
EXPOSE 8081
ENTRYPOINT ["/repo"]
