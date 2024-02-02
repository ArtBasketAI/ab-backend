FROM golang:1.21

ARG GITHUB_TOKEN
ARG SERVICE_NAME
ARG VERSION

WORKDIR /tmp/${SERVICE_NAME}

COPY go.* ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go install -v \
    -ldflags="-w -s -X main.version=${VERSION} -X main.serviceName=${SERVICE_NAME}" \
    ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=0 /go/bin/server /bin/server

EXPOSE 8080

CMD ["/bin/server"]
