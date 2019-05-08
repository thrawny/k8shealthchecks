FROM golang:1.12-alpine

RUN apk add --no-cache git build-base

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o k8shealthchecks

ENTRYPOINT ["./k8shealthchecks"]
