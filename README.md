<h1 style="text-align: center; color: aqua"> Reverse Proxy GO</h1>

    A reverse proxy that serves static files and targets based on defined routes

### Features:
* Serve static files
* Round Robin Load balancing for multiple target routes
* Retries different target if there's a downtime on a target

### How to run:
```bash
go build -o reverse-proxy-go main.go
./reverse-proxy-go config.json
```