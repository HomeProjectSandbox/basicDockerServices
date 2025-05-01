IMAGE_NAME=my-golang-app
CONTAINER_NAME=my-running-app

run:
	go run .
build:
	CGO_ENABLED=0 go build -a -o webserver main.go
build-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o webserver main.go
build-arm64:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -a -o webserver main.go

build_and_push:
	@if [ "$(shell uname -m)" = "arm64" ] || [ "$(shell uname -m)" = "aarch64" ]; then \
		docker build -t tamasujvari/basicsetup_basicwebserver:arm .; \
		docker push tamasujvari/basicsetup_basicwebserver:arm; \
	else \
		docker build -t tamasujvari/basicsetup_basicwebserver:latest .; \
		docker push tamasujvari/basicsetup_basicwebserver:latest; \
	fi

build_and_push_arm:
	docker buildx build -t tamasujvari/webserver_w_postgress_webserver:arm --platform linux/arm64 .
	docker push tamasujvari/webserver_w_postgress_webserver:arm

build_and_push_x86:
	docker buildx build -t tamasujvari/webserver_w_postgress_webserver:amd --platform linux/amd64 .
	docker push tamasujvari/webserver_w_postgress_webserver:amd