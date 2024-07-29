FROM golang:1.22-alpine
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY * ./
RUN CGO_ENABLED=0 go build -o shortener
CMD ["./shortener"]