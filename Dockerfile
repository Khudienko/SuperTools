
FROM golang:latest

WORKDIR /go/src/github.com/Khudienko/SuperTools

COPY . /go/src/github.com/Khudienko/SuperTools

RUN go build -o ./server ./cmd/main.go

CMD "./server"

EXPOSE 8080