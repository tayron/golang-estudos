VERSION=1.0.0
IMAGE-NAME=golangci-lint
CONTAINER-NAME=golangci-lint

GO-MAIN-SCRAPPER-NAME=main.go
BINARY-SCRAPPER-NAME=golangci

build-image:
	docker build -t ${IMAGE-NAME}:latest . -f Dockerfile

run-image:
	docker run -it --name ${CONTAINER-NAME} --rm ${IMAGE-NAME} sh

remove-image:
	docker image rmi -f ${IMAGE-NAME}

get-size-image:
	docker inspect -f "{{ .Size }}" ${IMAGE-NAME} | numfmt --to=si