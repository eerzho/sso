FROM golang:1.23-alpine

WORKDIR /http

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

CMD ["go", "run", "./cmd/http"]