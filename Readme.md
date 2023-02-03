# HTTP Request Limiter

This http request limiter were implemented as a middleware package - rlm.

# Default params:
 * Maximum Requests per second = 5 r/sec
 * Block time for too much request = 120 sec
 * Block time increment = 0 sec

# Ways to run project
Go to the project directory
```
git clone https://github.com/doxanocap/test-task-rlm.git
cd test-task-rlm

```
### Run this is you want to set up project as a docker container
```
docker compose up
```
### Run project manually  using
```
go run main.go
```
### If you want run project with following params
where:
 * A => input maximum requests per second (int)
 * B => input block time for too much request (int)
 * C => block time increment (int)
   
```
 go run main.go -i  A B C
```

# Run tests 
```
go test -v ./pkg/handler
```