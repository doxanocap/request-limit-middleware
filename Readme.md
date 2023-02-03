# HTTP Request Limiter

This http request limiter were implemented as a middleware package - rlm.

# Default params:
 * Maximum Requests per second = 5 r/sec
 * Block time for too much request = 120 sec
 * Block time increment = 0 sec

# Ways to run project
```
git clone https://github.com/doxanocap/test-task-rlm.git
cd test-task-rlm
 |
 v
 
docker compose up

or

go run main.go

#To run with params

go run main.go -i  A B C

where:
A => input maximum requests per second (int)
B => input block time for too much request (int)
C => block time increment (int)
```

# Run tests 
```
go test -v ./pkg/handler
```