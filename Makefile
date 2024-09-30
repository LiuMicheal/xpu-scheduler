.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags '-w' -o out/xpu-scheduler main.go

.PHONY: image #go build xpu-shceduler , build docker and cri image
image:
	sh make-image.sh

.PHONY: clear #clear docker image and cri image
clear:
	sh clear-image.sh