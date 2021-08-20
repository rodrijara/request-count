FROM golang:1.16-alpine

WORKDIR /server

COPY go.mod .
COPY src/main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o serverReqCount main.go

EXPOSE 8083

CMD [ "./serverReqCount" ]