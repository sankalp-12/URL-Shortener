# URL-Shortner


### Execution on localhost
 - Comment all the services in docker-compose.yml except api and db service. (remove network-mode: options from the api and db service)
 - Run docker-compose up (with wkdir: URL Shortner(Master))
 - Redis Exporter: 
	git clone https://github.com/oliver006/redis_exporter.git
	cd redis_exporter
	go build .
	./redis_exporter -redis.addr localhost:6379
 - Replace local prometheus.yml with /Prometheus/prometheus.yml & run ./prometheus
 
