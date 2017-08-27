.PHONY: default binary run container push

default: binary
binary: check-receiver

run:
	go run check-receiver.go

check-receiver: check-receiver.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o check-receiver .
	strip check-receiver

container: binary
	docker build -t mazdermind/check-receiver:latest .

push: container
	docker push mazdermind/check-receiver:latest
