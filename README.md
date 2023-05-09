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
- Docker image https://hub.docker.com/_/influxdb/tags
``` bash
# install influxdb server
docker pull influxdb:{tag}
```

## Test script

host.docker.internal = connect server outside container
``` javascript
// usage
http.get("http://host.docker.internal:5000/health")
```

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