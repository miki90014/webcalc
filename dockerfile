FROM golang:latest AS builder
ENV GOPROXY="direct"
ARG version
ARG gitCommit
ENV BUILD_VERSION=${version:-unknown}
ENV BUILD_GIT_COMMIT=${gitCommit:-unknown}
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN git config --global http.proxy http://135.245.48.34:8000
RUN http_proxy=http://135.245.48.34:8000/ https_proxy=http://135.245.48.34:8000/ go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o repo -ldflags "-X 'main.version=${BUILD_VERSION}' -X 'main.gitCommit=${BUILD_GIT_COMMIT}' -X 'main.binaryType=static' -w -extldflags '-static'" .

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/repo .
USER 8080:8080
USER 8081:8081
ENTRYPOINT ["/repo"]
