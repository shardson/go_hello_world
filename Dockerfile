FROM golang:1.22-alpine

WORKDIR /

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN go build -o main.go

EXPOSE 8080

CMD ["./main.go"]