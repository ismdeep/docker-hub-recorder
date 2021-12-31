SHELL=/bin/bash

help:

build-vendor:
	go mod tidy
	go mod vendor

run:
	go run main.go

test-renew:
	-docker stop docker-hub-recorder-db
	-docker rm   docker-hub-recorder-db
	cp $(CURDIR)/.data/config.local.yaml $(CURDIR)/config.yaml
	docker run --name docker-hub-recorder-db \
		-p 10004:3306 \
		-e MYSQL_ROOT_PASSWORD=docker-hub-recorder-123456 \
		-e MYSQL_DATABASE=docker_hub_recorder \
		-d hub.deepin.com/library/mysql:8.0

test-clean:
	-docker stop docker-hub-recorder-db
	-docker rm   docker-hub-recorder-db