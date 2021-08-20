# Golang Developer Test

## To do

1. "RequestCount" server:

   - counts the number of requests handled by the server instance
   - counts the total number of requests handled by all servers that are running
   - serves these numbers through HTTP, with text/plain response content

2. Dockerize the "RequestCount" server

   - tip: build the binary with CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main_name" main_name.go

3. Deploy multiple instances of the "RequestCount" server with Docker Compose (replicas: 3)

Example flow:

```shell
$ curl http://localhost:8083
You are talking to instance host2:8083.
This is request 1 to this instance and request 1 to the cluster.

$ curl http://localhost:8083
You are talking to instance host1:8083.
This is request 1 to this instance and request 2 to the cluster.

$ curl http://localhost:8083
You are talking to instance host2:8083.
This is request 2 to this instance and request 3 to the cluster.
```

# Solution

## Usage

Clone this repository and make sure to have [Docker Engine](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed and running.

1. Once inside `request-count` (cloned repo), run:

```shell
$ docker-compose up
```

2. If you see `ready for start up` at the last line, it means the server is running.

```shell
# skipping lines here
  nginx_1 | /docker-entrypoint.sh: Configuration complete; ready for start up
```

3. Make requests with curl and see how many times you visited one of the instances and also how many times in total you visited the cluster.

```shell
$ curl http://localhost:8083
```

## 1. "RequestCount" server:

The number of visit to the server has been calculated using log files.

`visitHandler` function runs when a request to the server `localhost:8083` is made.

`WriteLogs` creates/open files with instance (replica) and cluster names and writes a line
with datetime information.

`VisitCounts` takes as input the paths of replica/container and cluster log files and returns the counts of the lines written both of them. The counts are captured and used in the response message.
Instance name also is used in the response message, so that the count of visits is referenced.

```golang
func visitHandler(res http.ResponseWriter, req *http.Request) {

	instance, err := os.Hostname()
	if err != nil {
		log.Println("ERROR getting instance Hostname", err)
	}
	cluster := req.Host

	WriteLogs(DBDIR+instance, DBDIR+cluster)

	icount, ccount := VisitCounts(DBDIR+instance, DBDIR+cluster)

	fmt.Fprintf(res,
		"You are talking to instance %s:%s.\nThis is request %d to this instance and request %d to the cluster.\n",
		instance, PORT, icount, ccount)
}
```

## 2. Dockerize the "RequestCount" server

Dockerfile specifies a single-stage build from a golang alpine image, so the final image is much lighter.\
Port `8083` is exposed.

```dockerfile
FROM golang:1.16-alpine

WORKDIR /myapp

COPY go.mod .
COPY src/main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /requestCount main.go

EXPOSE 8083

CMD [ "/requestCount" ]
```

## 3. Deploy multiple instances of the "RequestCount" server with Docker Compose (replicas: 3)

To deploy three replicas of the server container that can be accessed from a single call to `localhost:8083`, a nginx reverse proxy server was set. The official last image was used.

Since the containers need to write and read log files, and the cluster log files must be shared between them, a named volume `mydb` was set.

```yaml
version: "3.9"
services:
  reqserver:
    build: .
    deploy:
      replicas: 3
    volumes:
      - mydb:/db
  nginx:
    image: nginx:latest
    ports:
      - "8083:8083"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro

volumes:
  mydb:
```
