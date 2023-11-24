FROM golang:1.17-alpine as builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY main.go .

RUN go build -v -o /usr/src/app/go-crud-rest-api

FROM alpine:3.18

COPY --from=builder /usr/src/app/go-crud-rest-api /app/go-crud-rest-api

CMD ["/app/go-crud-rest-api"]
