# Modules Caching
FROM golang:latest as modules
COPY go.sum go.mod /modules/
WORKDIR /modules
RUN go mod download

# Builder
FROM golang:latest as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

CMD ["go", "run", "."]
