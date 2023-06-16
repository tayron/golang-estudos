VERSION=1.0.0
IMAGE-NAME=scrapper-go
CONTAINER-NAME=scrapper-go

GO-MAIN-SCRAPPER-NAME=main.go
BINARY-SCRAPPER-NAME=scrapper

build-scrapper:
	go build -o ${BINARY-SCRAPPER-NAME} -ldflags "-s -w" ${GO-MAIN-SCRAPPER-NAME} 
	upx ${BINARY-SCRAPPER-NAME}

build-image:
	docker build -t ${IMAGE-NAME}:latest . -f Dockerfile

run-image:
	docker run -it --name ${CONTAINER-NAME} --rm ${IMAGE-NAME} sh

remove-image:
	docker image rmi -f ${IMAGE-NAME}

get-size-image:
	docker inspect -f "{{ .Size }}" ${IMAGE-NAME} | numfmt --to=si