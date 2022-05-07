# Simple web service for performance test

- run `go run main.go` to start

## Docker version

- you can get the image using this command `docker pull fabiansyah/pt-server`

---

## Monitoring information

### How to use

- change directory to monitoring
- run `docker-compose up -d`

### Grafana setting

- go to `http://localhost:3000/` to access grafana
- go to configuration and add new data source
- pick prometheus
- for prometheus HTTP URL use `http://km-prometheus:9090`
- set scrape interval to `3s`
- create new panel and use this metric browser

  ```promql
  sum(rate(http_request_count{}[5s])) by (response_code)
  ```

- if nothing show up on the panel, change `5s` to something higher like `6s-10s`
- don't forget to change time range to `5 minutes`

### Note

- check firewall rule or disable it so host.docker.internal will work
- scrape interval for prometheus (on prom and grafana)
