FROM golang:1.24

WORKDIR /app
COPY . .
ENV GOPROXY=direct
RUN go mod tidy
RUN go build -o server ./cmd/server

CMD ["./server"]
