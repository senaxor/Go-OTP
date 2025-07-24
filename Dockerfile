FROM golang:1.24

WORKDIR /app
COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download

COPY . .

# The source code will be mounted at runtime, so no COPY . .
CMD ["go", "run", "./cmd/server/main.go"]
