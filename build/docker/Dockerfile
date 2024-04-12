# 0. Fetch ca-certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# 1. build executable
FROM golang:1.22-alpine as builder

# 2. Add user for non root access
RUN addgroup -g 2000 appuser && \
    adduser -D -H -u 2000 -G appuser appuser

# 2. install build dependencies
RUN apk add git
RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# 3. generate proto files
WORKDIR /go/src/app
COPY protobuf/*.proto protobuf/
RUN protoc \
  --go_out=. \
  --go-grpc_out=. \
  protobuf/*.proto

# 4. install dependencies
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...
RUN go install -v ./...

# 5. Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o /tmp/main

# 5. Start process from scratch as appuser
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /tmp/main /bin/main
USER appuser

ENTRYPOINT [ "/bin/main" ]