FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest


# Build the Go app
# RUN go build -o main .

EXPOSE 8080

CMD ["air"]