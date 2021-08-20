# Golang Developer Test

# To do

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
