FROM golang:1.22-alpine
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 go build -C main -o shortener
RUN chmod +x ./main/shortener
CMD ["./main/shortener"]