build-scrapper:
	go build -o scrapper -ldflags "-s -w" main.go && upx scrapper

build-docker-image:
	docker build -t scrapper-go:latest . -f Dockerfile	

run:
	docker run -it --name scrapper-go --rm scrapper-go sh