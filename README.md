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

# Solution

1. "RequestCount" server:

   Visit counts to the server has been calculated from log files.
   `visitHandler` function runs when a request to the server `localhost:8083` is made.\
   `WriteLogs` creates/open files with instance (replica cointainer) and cluster names and writes a line
   with datetime information.
   `VisitCounts` takes as input the paths of replica container and cluster log files and returns the counts of the lines written both of them. The counts are captured and used in the response message.
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
