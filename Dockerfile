FROM golang:latest as builder

WORKDIR /go/app/

COPY go.* /go/app/

RUN go mod download

COPY . .

RUN mkdir -p bin

RUN GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/ ./cmd/...

FROM alpine:3.12

RUN mkdir /app/

COPY --from=builder /go/app/bin /app/

WORKDIR /app/