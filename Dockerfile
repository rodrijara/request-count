FROM golang:1.16-alpine

WORKDIR /myapp 

COPY go.mod .
COPY src/*.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /requestCount main.go functions.go

EXPOSE 8083

CMD [ "/requestCount" ]