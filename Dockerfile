FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 4000
EXPOSE 587

CMD ["go", "run", "./cmd/api"]