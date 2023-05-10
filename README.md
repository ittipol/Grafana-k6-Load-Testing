# Grafana k6 - Load testing
- https://k6.io/

## Grafana k6 install
``` bash
# use brew
brew install k6

# use docker
docker pull loadimpact/k6

# run script
docker compose run --rm k6 run /scripts/test.js

# help
docker compose run --rm k6 run --help
```

## InfluxDB
- Document https://docs.influxdata.com/influxdb
- InfluxDB environment variable https://docs.influxdata.com/influxdb/v1.8/administration/config/
- Docker image https://hub.docker.com/_/influxdb/tags
``` bash
# install influxdb server
docker pull influxdb:{tag}
```

## Grafana setup & configuration
## Grafana Dashboard
### Add data source from InfluxDB
1. Go to http://localhost:3000/
2. Click menu "Administration" > "Data sources"
3. Click button "Add new data source"
4. Click "InfluxDB"
5. Set Name 
6. Set Url = http://influxdb:8086
7. Set Database = k6
8. Click "Save & test"

### Search dashboard k6 ID
1. Go to https://grafana.com/grafana/dashboards/
2. Type "k6" in Search dashboards input box
3. k6 Load Testing Results https://grafana.com/grafana/dashboards/2587-k6-load-testing-results/
4. Click button "Copy ID to clipboard"

### Import dashboard
1. Go to http://localhost:3000/
2. Click menu "Dashboard" > "New" > "Import"
3. Import via grafana.com Input "ID"
4. Click "Load"
5. Select a InfluxDB data source > select InfluxDB
6. Click "Import"

## Redis configuration
- Download configuration file https://redis.io/docs/management/config/
``` bash
# NETWORK
bind 0.0.0.0

# Use APPEND ONLY MODE
appendonly yes

# Disable SNAPSHOTTING
save ""
```

## Redis command
- https://redis.io/commands/

``` bash
AUTH

CLIENT CACHING

CLIENT GETNAME

CLIENT GETREDIR

CLIENT ID

CLIENT INFO

CLIENT KILL

CLIENT LIST

CLIENT NO-EVICT

CLIENT NO-TOUCH

CLIENT PAUSE

CLIENT REPLY

CLIENT SETINFO

CLIENT SETNAME

CLIENT TRACKING

CLIENT TRACKINGINFO

CLIENT UNBLOCK

CLIENT UNPAUSE

ECHO

HELLO

PING

QUIT

RESET

SELECT

MONITOR

FLUSHALL

FLUSHDB

INFO
```

## Redis client for Go
- https://redis.io/resources/clients/#go
### To install go-redis/v9:
``` bash
go get github.com/redis/go-redis/v9
```
### To connect to a Redis server:
``` go
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

client := redis.NewClient(&redis.Options{
	Addr:	  "localhost:6379",
	Password: "", // no password set
	DB:		  0,  // use default DB
})
```

### Another way to connect is using a connection string.
``` go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

client := redis.NewClient(opt)
```

## k6 test script
k6 use javascript ES6 to create test script
``` javascript
// usage
import http from 'k6/http'

export let options = {
    vus: 5,
    duration: '10s'
}

export default function() {
    // host.docker.internal = connect service outside container
    http.get("http://host.docker.internal:5000/health")
}
```

## Go Packages
- redis [https://pkg.go.dev/github.com/redis/go-redis/v9](https://pkg.go.dev/github.com/redis/go-redis/v9)